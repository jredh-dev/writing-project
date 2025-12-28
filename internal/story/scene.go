package story

import "fmt"

// Scene represents a single story beat
type Scene struct {
	ID         string
	ThreadType ThreadType
	Text       string
	Choices    []Choice
	Next       string // For open/affirmative/finisher thread types
	MinLength  int    // For open responses
}

// Choice represents an option the player can select
type Choice struct {
	Text   string
	Next   string // Scene ID to transition to
	Impact string // Format: "entity.attribute±value" (stored, not processed yet)
}

// ValidationResult is returned after validating a response
type ValidationResult struct {
	Valid       bool
	Feedback    string
	NextSceneID string
}

// Global cache for loaded scenes
var sceneCache []Scene
var sceneMap map[string]*Scene

// GetScene returns a scene by ID
func GetScene(id string) *Scene {
	if sceneMap == nil {
		loadScenes()
	}
	return sceneMap[id]
}

// GetPrefaceScenes returns all preface scenes loaded from YAML
// Edit scenes/preface.yaml to write your story!
func GetPrefaceScenes() []Scene {
	if len(sceneCache) > 0 {
		return sceneCache
	}
	loadScenes()
	return sceneCache
}

// loadScenes loads and caches scenes from YAML
func loadScenes() {
	scenes, err := LoadScenesFromYAML("scenes/preface.yaml")
	if err != nil {
		// In production, this should panic or return error
		// For now, return empty to allow graceful degradation
		panic(fmt.Sprintf("Failed to load scenes: %v", err))
	}

	sceneCache = scenes

	// Build scene map for quick lookup
	sceneMap = make(map[string]*Scene)
	for i := range sceneCache {
		sceneMap[sceneCache[i].ID] = &sceneCache[i]
	}
}

// LoadScenes allows explicitly loading scenes (useful for testing)
func LoadScenes(scenes []Scene) {
	sceneCache = scenes
	sceneMap = make(map[string]*Scene)
	for i := range sceneCache {
		sceneMap[sceneCache[i].ID] = &sceneCache[i]
	}
}
