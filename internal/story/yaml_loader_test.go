package story

import (
	"strings"
	"testing"
)

func TestLoadScenesFromYAML(t *testing.T) {
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}

	if len(scenes) == 0 {
		t.Fatal("Expected at least one scene to be loaded")
	}

	// Test first scene (preface.0:dream-start)
	dreamStart := scenes[0]
	if dreamStart.ID != "preface.0:dream-start" {
		t.Errorf("Expected first scene ID to be 'preface.0:dream-start', got '%s'", dreamStart.ID)
	}

	if dreamStart.ThreadType != ThreadMulti {
		t.Errorf("Expected thread_type 'multi', got '%s'", dreamStart.ThreadType)
	}

	if len(dreamStart.Choices) != 3 {
		t.Errorf("Expected 3 choices for dream-start, got %d", len(dreamStart.Choices))
	}

	// Test that text was loaded
	if dreamStart.Text == "" {
		t.Error("Expected scene text to be loaded")
	}

	// Test first choice has impact
	if dreamStart.Choices[0].Impact == "" {
		t.Error("Expected first choice to have impact")
	}
	if !strings.Contains(dreamStart.Choices[0].Impact, "player.strength") {
		t.Errorf("Expected impact to contain 'player.strength', got: %s", dreamStart.Choices[0].Impact)
	}
}

func TestGetScene(t *testing.T) {
	// Load scenes first
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}
	LoadScenes(scenes)

	// Test GetScene function
	scene := GetScene("preface.0:dream-start")
	if scene == nil {
		t.Fatal("Expected to find 'preface.0:dream-start' scene")
	}

	if scene.ID != "preface.0:dream-start" {
		t.Errorf("Expected scene ID 'preface.0:dream-start', got '%s'", scene.ID)
	}

	// Test non-existent scene
	nonExistent := GetScene("does-not-exist")
	if nonExistent != nil {
		t.Error("Expected nil for non-existent scene")
	}
}

func TestOpenResponseScene(t *testing.T) {
	// Load scenes
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}

	// Find teacher-choice scene
	var teacherChoice *Scene
	for i := range scenes {
		if scenes[i].ID == "preface.3:teacher-choice" {
			teacherChoice = &scenes[i]
			break
		}
	}

	if teacherChoice == nil {
		t.Fatal("Expected to find 'preface.3:teacher-choice' scene")
	}

	if teacherChoice.ThreadType != ThreadOpen {
		t.Errorf("Expected thread_type 'open', got '%s'", teacherChoice.ThreadType)
	}

	if teacherChoice.MinLength != 10 {
		t.Errorf("Expected min_length 10, got %d", teacherChoice.MinLength)
	}

	if teacherChoice.Next != "preface.4:assigned-teacher" {
		t.Errorf("Expected next 'preface.4:assigned-teacher', got '%s'", teacherChoice.Next)
	}
}

func TestSceneProgression(t *testing.T) {
	// Load scenes first
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}
	LoadScenes(scenes)

	// Test that all scenes in the progression exist
	expectedScenes := []string{
		"preface.0:dream-start",
		"preface.1:registration",
		"preface.2:campus-tour",
		"preface.3:teacher-choice",
		"preface.4:assigned-teacher",
		"preface.5:tutorial-multiple",
		"preface.6:end-of-demo",
	}

	for _, sceneID := range expectedScenes {
		scene := GetScene(sceneID)
		if scene == nil {
			t.Errorf("Expected scene '%s' to exist", sceneID)
		}
	}
}

func TestGraphValidation(t *testing.T) {
	// Graph validation happens automatically in LoadScenesFromYAML
	_, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Graph validation failed: %v", err)
	}
}

func TestThreadTypes(t *testing.T) {
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}

	// Test that each thread type is represented
	threadTypes := map[ThreadType]bool{}
	for _, scene := range scenes {
		threadTypes[scene.ThreadType] = true
	}

	expectedTypes := []ThreadType{ThreadMulti, ThreadOpen, ThreadAffirmative}
	for _, expectedType := range expectedTypes {
		if !threadTypes[expectedType] {
			t.Errorf("Expected to find at least one scene with thread_type '%s'", expectedType)
		}
	}
}

func TestTerminalScene(t *testing.T) {
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}

	// Find end-of-demo scene
	var endScene *Scene
	for i := range scenes {
		if scenes[i].ID == "preface.6:end-of-demo" {
			endScene = &scenes[i]
			break
		}
	}

	if endScene == nil {
		t.Fatal("Expected to find 'preface.6:end-of-demo' scene")
	}

	if endScene.Next != "0" {
		t.Errorf("Expected terminal scene to have next='0', got '%s'", endScene.Next)
	}
}

func TestImpactFormat(t *testing.T) {
	scenes, err := LoadScenesFromYAML("../../scenes/preface.yaml")
	if err != nil {
		t.Fatalf("Failed to load scenes: %v", err)
	}

	// Check that impacts are properly formatted
	dreamStart := scenes[0]
	for i, choice := range dreamStart.Choices {
		if choice.Impact == "" {
			t.Errorf("Choice %d expected to have impact", i)
		}

		// Should match format: entity.attribute±value
		if !strings.Contains(choice.Impact, ".") ||
			(!strings.Contains(choice.Impact, "+") && !strings.Contains(choice.Impact, "-")) {
			t.Errorf("Choice %d impact '%s' doesn't match format entity.attribute±value", i, choice.Impact)
		}
	}
}
