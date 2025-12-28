package story

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// ThreadType defines how a scene presents choices to the player
type ThreadType string

const (
	ThreadMulti       ThreadType = "multi"       // Multiple choice with different next scenes
	ThreadOpen        ThreadType = "open"        // Text input with validation
	ThreadAffirmative ThreadType = "affirmative" // Simple "Continue" button
	ThreadFinisher    ThreadType = "finisher"    // (Stubbed) Real-time tokenization
)

// YAMLScene represents a scene as defined in YAML
type YAMLScene struct {
	ID         string       `yaml:"id"`
	ThreadType ThreadType   `yaml:"thread_type"`
	Text       string       `yaml:"text"`
	Choices    []YAMLChoice `yaml:"choices,omitempty"`
	Validation *struct {
		MinLength int `yaml:"min_length"`
	} `yaml:"validation,omitempty"`
	Next string `yaml:"next,omitempty"` // For open/affirmative/finisher
}

// YAMLChoice represents a choice option in YAML
type YAMLChoice struct {
	Text   string `yaml:"text"`
	Next   string `yaml:"next"`
	Impact string `yaml:"impact,omitempty"` // Format: "entity.attribute±value"
}

// YAMLSceneFile represents the top-level YAML structure
type YAMLSceneFile struct {
	Scenes []YAMLScene `yaml:"scenes"`
}

// LoadScenesFromYAML loads scenes from a YAML file and validates the graph
func LoadScenesFromYAML(filename string) ([]Scene, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var sceneFile YAMLSceneFile
	if err := yaml.Unmarshal(data, &sceneFile); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Convert YAML scenes to Scene structs
	scenes := make([]Scene, 0, len(sceneFile.Scenes))
	sceneMap := make(map[string]*Scene) // For quick lookup during validation

	for _, yamlScene := range sceneFile.Scenes {
		scene, err := convertYAMLScene(yamlScene)
		if err != nil {
			return nil, fmt.Errorf("failed to convert scene %s: %w", yamlScene.ID, err)
		}
		scenes = append(scenes, scene)
		sceneMap[scene.ID] = &scenes[len(scenes)-1]
	}

	// Validate the scene graph
	if err := validateSceneGraph(scenes, sceneMap); err != nil {
		return nil, fmt.Errorf("scene graph validation failed: %w", err)
	}

	return scenes, nil
}

// convertYAMLScene converts a YAMLScene to a Scene
func convertYAMLScene(yamlScene YAMLScene) (Scene, error) {
	// Validate ID format: chapter.scene-number:description
	if err := validateSceneID(yamlScene.ID); err != nil {
		return Scene{}, err
	}

	// Convert choices
	choices := make([]Choice, 0, len(yamlScene.Choices))
	for _, yamlChoice := range yamlScene.Choices {
		choices = append(choices, Choice{
			Text:   yamlChoice.Text,
			Next:   yamlChoice.Next,
			Impact: yamlChoice.Impact,
		})
	}

	// Create scene
	scene := Scene{
		ID:         yamlScene.ID,
		ThreadType: yamlScene.ThreadType,
		Text:       strings.TrimSpace(yamlScene.Text),
		Choices:    choices,
		Next:       yamlScene.Next, // For open/affirmative/finisher
	}

	// Add validation for open responses
	if yamlScene.ThreadType == ThreadOpen && yamlScene.Validation != nil {
		scene.MinLength = yamlScene.Validation.MinLength
	}

	return scene, nil
}

// validateSceneID validates the scene ID format: chapter.scene-number:description
func validateSceneID(id string) error {
	if id == "0" {
		// Special case: terminal scene
		return nil
	}

	// Pattern: chapter.number:description
	// Examples: preface.0:dream-start, chapter1.5:boss-fight
	pattern := `^[a-z0-9]+\.[0-9]+(\.[0-9]+)?:[a-z0-9-]+$`
	matched, err := regexp.MatchString(pattern, id)
	if err != nil {
		return fmt.Errorf("regex error: %w", err)
	}
	if !matched {
		return fmt.Errorf("invalid scene ID format '%s': must be chapter.number:description (e.g., preface.0:dream-start)", id)
	}
	return nil
}

// validateSceneGraph validates the scene graph structure
func validateSceneGraph(scenes []Scene, sceneMap map[string]*Scene) error {
	errors := []string{}

	for _, scene := range scenes {
		// Check thread type
		switch scene.ThreadType {
		case ThreadMulti:
			// Multi must have choices
			if len(scene.Choices) == 0 {
				errors = append(errors, fmt.Sprintf("scene %s: thread_type 'multi' requires at least one choice", scene.ID))
			}
			// Each choice must have a valid 'next'
			for i, choice := range scene.Choices {
				if err := validateNext(choice.Next, scene.ID, sceneMap); err != nil {
					errors = append(errors, fmt.Sprintf("scene %s choice %d: %v", scene.ID, i, err))
				}
			}

		case ThreadOpen, ThreadAffirmative, ThreadFinisher:
			// These must have scene-level 'next'
			if scene.Next == "" {
				errors = append(errors, fmt.Sprintf("scene %s: thread_type '%s' requires 'next' field at scene level", scene.ID, scene.ThreadType))
			}
			if err := validateNext(scene.Next, scene.ID, sceneMap); err != nil {
				errors = append(errors, fmt.Sprintf("scene %s: %v", scene.ID, err))
			}

		default:
			errors = append(errors, fmt.Sprintf("scene %s: invalid thread_type '%s' (must be multi, open, affirmative, or finisher)", scene.ID, scene.ThreadType))
		}

		// Validate impact format if present
		for i, choice := range scene.Choices {
			if choice.Impact != "" {
				if err := validateImpact(choice.Impact); err != nil {
					errors = append(errors, fmt.Sprintf("scene %s choice %d: invalid impact format: %v", scene.ID, i, err))
				}
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("\n  - %s", strings.Join(errors, "\n  - "))
	}

	return nil
}

// validateNext validates that a 'next' value is valid
func validateNext(next string, currentSceneID string, sceneMap map[string]*Scene) error {
	if next == "0" {
		// Terminal scene (exit successfully)
		return nil
	}

	if strings.HasPrefix(next, "-") {
		// Negative values reserved for error states (valid for now)
		return nil
	}

	// Check if the scene exists
	if _, exists := sceneMap[next]; !exists {
		return fmt.Errorf("next '%s' references non-existent scene", next)
	}

	return nil
}

// validateImpact validates impact string format: entity.attribute±value
func validateImpact(impact string) error {
	// Pattern: entity.attribute(.subattribute)*±value
	// Examples: player.strength+2, npc.teacher.trust-5, world.chaos+10
	pattern := `^[a-z_]+(\.[a-z_]+)*[+-]\d+$`
	matched, err := regexp.MatchString(pattern, impact)
	if err != nil {
		return fmt.Errorf("regex error: %w", err)
	}
	if !matched {
		return fmt.Errorf("'%s' must be format entity.attribute±value (e.g., player.strength+2)", impact)
	}
	return nil
}
