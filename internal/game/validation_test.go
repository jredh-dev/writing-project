package game

import (
	"testing"
)

func TestResponseValidator_Validate(t *testing.T) {
	tests := []struct {
		name             string
		keywords         []string
		requiredCount    int
		input            string
		wantCorrect      bool
		wantMatchedCount int
		wantScore        float64
	}{
		{
			name:             "exact match single keyword",
			keywords:         []string{"1914"},
			requiredCount:    1,
			input:            "1914",
			wantCorrect:      true,
			wantMatchedCount: 1,
			wantScore:        1.0,
		},
		{
			name:             "keyword in sentence",
			keywords:         []string{"division"},
			requiredCount:    1,
			input:            "I think the answer is division",
			wantCorrect:      true,
			wantMatchedCount: 1,
			wantScore:        1.0,
		},
		{
			name:             "multiple keywords all present",
			keywords:         []string{"math", "division", "algebra"},
			requiredCount:    3,
			input:            "Math uses division and algebra",
			wantCorrect:      true,
			wantMatchedCount: 3,
			wantScore:        1.0,
		},
		{
			name:             "partial match - some keywords missing",
			keywords:         []string{"math", "division", "algebra"},
			requiredCount:    3,
			input:            "Math uses division",
			wantCorrect:      false,
			wantMatchedCount: 2,
			wantScore:        0.666,
		},
		{
			name:             "case insensitive matching",
			keywords:         []string{"division"},
			requiredCount:    1,
			input:            "DIVISION",
			wantCorrect:      true,
			wantMatchedCount: 1,
			wantScore:        1.0,
		},
		{
			name:             "punctuation handling",
			keywords:         []string{"division"},
			requiredCount:    1,
			input:            "The answer is: division.",
			wantCorrect:      true,
			wantMatchedCount: 1,
			wantScore:        1.0,
		},
		{
			name:             "no keywords matched",
			keywords:         []string{"division"},
			requiredCount:    1,
			input:            "I don't know",
			wantCorrect:      false,
			wantMatchedCount: 0,
			wantScore:        0.0,
		},
		{
			name:             "required count less than total",
			keywords:         []string{"math", "division", "algebra"},
			requiredCount:    2,
			input:            "Math and division",
			wantCorrect:      true,
			wantMatchedCount: 2,
			wantScore:        0.666,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator(tt.keywords, tt.requiredCount)
			result := v.Validate(tt.input)

			if result.Correct != tt.wantCorrect {
				t.Errorf("Correct = %v, want %v", result.Correct, tt.wantCorrect)
			}

			if len(result.MatchedWords) != tt.wantMatchedCount {
				t.Errorf("MatchedWords count = %d, want %d", len(result.MatchedWords), tt.wantMatchedCount)
			}

			// Allow small floating point difference
			if abs(result.Score-tt.wantScore) > 0.01 {
				t.Errorf("Score = %.3f, want %.3f", result.Score, tt.wantScore)
			}
		})
	}
}

func TestResponseValidator_AnnotateMatches(t *testing.T) {
	v := NewValidator([]string{"division", "math"}, 1)
	result := v.Validate("I learned division and math today")

	annotated := string(result.AnnotatedInput)

	if !containsString(annotated, `<mark class="match-correct">division</mark>`) {
		t.Errorf("Expected 'division' to be highlighted, got: %s", annotated)
	}

	if !containsString(annotated, `<mark class="match-correct">math</mark>`) {
		t.Errorf("Expected 'math' to be highlighted, got: %s", annotated)
	}
}

func TestNumericValidator_ValidateNumeric(t *testing.T) {
	tests := []struct {
		name        string
		accepted    []float64
		tolerance   float64
		input       string
		wantCorrect bool
	}{
		{
			name:        "exact match",
			accepted:    []float64{1914},
			tolerance:   0,
			input:       "1914",
			wantCorrect: true,
		},
		{
			name:        "within tolerance",
			accepted:    []float64{3.14},
			tolerance:   0.01,
			input:       "3.15",
			wantCorrect: true,
		},
		{
			name:        "outside tolerance",
			accepted:    []float64{3.14},
			tolerance:   0.01,
			input:       "3.2",
			wantCorrect: false,
		},
		{
			name:        "multiple accepted values",
			accepted:    []float64{1914, 1915},
			tolerance:   0,
			input:       "1915",
			wantCorrect: true,
		},
		{
			name:        "invalid input",
			accepted:    []float64{1914},
			tolerance:   0,
			input:       "not a number",
			wantCorrect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NumericValidator{
				AcceptedValues: tt.accepted,
				Tolerance:      tt.tolerance,
			}
			result := v.ValidateNumeric(tt.input)

			if result.Correct != tt.wantCorrect {
				t.Errorf("Correct = %v, want %v", result.Correct, tt.wantCorrect)
			}
		})
	}
}

func TestMultipleChoiceValidator_ValidateChoice(t *testing.T) {
	tests := []struct {
		name           string
		correctIndices []int
		allowMultiple  bool
		selected       []int
		wantCorrect    bool
	}{
		{
			name:           "single correct choice",
			correctIndices: []int{2},
			allowMultiple:  false,
			selected:       []int{2},
			wantCorrect:    true,
		},
		{
			name:           "single wrong choice",
			correctIndices: []int{2},
			allowMultiple:  false,
			selected:       []int{1},
			wantCorrect:    false,
		},
		{
			name:           "multiple selections not allowed",
			correctIndices: []int{2},
			allowMultiple:  false,
			selected:       []int{1, 2},
			wantCorrect:    false,
		},
		{
			name:           "all correct choices selected",
			correctIndices: []int{1, 3},
			allowMultiple:  true,
			selected:       []int{1, 3},
			wantCorrect:    true,
		},
		{
			name:           "missing one correct choice",
			correctIndices: []int{1, 3},
			allowMultiple:  true,
			selected:       []int{1},
			wantCorrect:    false,
		},
		{
			name:           "includes wrong choice",
			correctIndices: []int{1, 3},
			allowMultiple:  true,
			selected:       []int{1, 2, 3},
			wantCorrect:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &MultipleChoiceValidator{
				CorrectIndices: tt.correctIndices,
				AllowMultiple:  tt.allowMultiple,
			}
			result := v.ValidateChoice(tt.selected)

			if result.Correct != tt.wantCorrect {
				t.Errorf("Correct = %v, want %v", result.Correct, tt.wantCorrect)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		caseSensitive bool
		want          string
	}{
		{
			name:          "trim whitespace",
			input:         "  hello  ",
			caseSensitive: false,
			want:          "hello",
		},
		{
			name:          "collapse multiple spaces",
			input:         "hello    world",
			caseSensitive: false,
			want:          "hello world",
		},
		{
			name:          "case insensitive",
			input:         "HELLO",
			caseSensitive: false,
			want:          "hello",
		},
		{
			name:          "case sensitive",
			input:         "HELLO",
			caseSensitive: true,
			want:          "HELLO",
		},
		{
			name:          "remove trailing punctuation",
			input:         "hello.",
			caseSensitive: false,
			want:          "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(tt.input, tt.caseSensitive)
			if got != tt.want {
				t.Errorf("normalize() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Helper functions
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func containsString(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && s != substr && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
