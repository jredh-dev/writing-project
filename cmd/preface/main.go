package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jredh-dev/divine-academy/internal/story"
)

type PageData struct {
	Scene    *story.Scene
	Feedback string
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("web/templates/*.html"))
}

func main() {
	// Load scenes on startup (will panic if validation fails)
	story.GetPrefaceScenes()
	fmt.Println("✅ Scene graph validated successfully")

	// Serve static files (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Main routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/scene", handleScene)
	http.HandleFunc("/choice", handleChoice)

	port := ":8080"
	fmt.Printf("\n🎮 Writing Project Preface running at http://localhost%s\n\n", port)
	fmt.Println("Open your browser and visit the URL above to play!")
	fmt.Println("Press Ctrl+C to stop the server.")
	fmt.Println()

	log.Fatal(http.ListenAndServe(port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Start with the first scene
	scene := story.GetScene("preface.0:dream-start")
	if scene == nil {
		http.Error(w, "Starting scene not found", http.StatusNotFound)
		return
	}

	renderScene(w, scene, "")
}

func handleScene(w http.ResponseWriter, r *http.Request) {
	sceneID := r.URL.Query().Get("id")
	if sceneID == "" {
		http.Error(w, "Missing scene ID", http.StatusBadRequest)
		return
	}

	scene := story.GetScene(sceneID)
	if scene == nil {
		http.Error(w, "Scene not found", http.StatusNotFound)
		return
	}

	renderScene(w, scene, "")
}

func handleChoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	sceneID := r.FormValue("scene_id")
	choiceIndexStr := r.FormValue("choice_index")
	userText := r.FormValue("user_text") // For open responses

	currentScene := story.GetScene(sceneID)
	if currentScene == nil {
		http.Error(w, "Current scene not found", http.StatusNotFound)
		return
	}

	var nextSceneID string
	var feedback string

	switch currentScene.ThreadType {
	case story.ThreadMulti:
		// Get the choice by index
		choiceIndex, err := strconv.Atoi(choiceIndexStr)
		if err != nil || choiceIndex < 0 || choiceIndex >= len(currentScene.Choices) {
			http.Error(w, "Invalid choice", http.StatusBadRequest)
			return
		}

		choice := currentScene.Choices[choiceIndex]
		nextSceneID = choice.Next
		feedback = fmt.Sprintf("You chose: %s", choice.Text)

		// TODO: Process impact (choice.Impact) when attribute system is implemented

	case story.ThreadOpen:
		// Validate open response
		if len(userText) < currentScene.MinLength {
			// Re-render current scene with error
			renderScene(w, currentScene, fmt.Sprintf("Please provide at least %d characters.", currentScene.MinLength))
			return
		}
		nextSceneID = currentScene.Next
		feedback = "Response recorded."

	case story.ThreadAffirmative, story.ThreadFinisher:
		// Simple continue
		nextSceneID = currentScene.Next
		feedback = ""

	default:
		http.Error(w, "Invalid thread type", http.StatusInternalServerError)
		return
	}

	// Check for terminal scene
	if nextSceneID == "0" {
		renderEndScreen(w)
		return
	}

	// Get next scene
	nextScene := story.GetScene(nextSceneID)
	if nextScene == nil {
		http.Error(w, "Next scene not found", http.StatusNotFound)
		return
	}

	renderScene(w, nextScene, feedback)
}

func renderScene(w http.ResponseWriter, scene *story.Scene, feedback string) {
	data := PageData{
		Scene:    scene,
		Feedback: feedback,
	}

	if err := templates.ExecuteTemplate(w, "scene.html", data); err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func renderEndScreen(w http.ResponseWriter) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Demo Complete</title>
		<link rel="stylesheet" href="/static/css/main.css">
	</head>
	<body>
		<main class="scene-container">
			<header><h1>Writing Project: Preface</h1></header>
			<article class="scene">
				<h2>Demo Complete!</h2>
				<p>You've reached the end of the current demo.</p>
				<p>The full preface will have many more scenes!</p>
				<br>
				<a href="/" style="color: #667eea; font-weight: bold;">Start Over</a>
			</article>
		</main>
	</body>
	</html>
	`)
}
