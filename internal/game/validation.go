package game

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"unicode"
)

// ResponseValidator checks free-text answers for open response questions
type ResponseValidator struct {
	AcceptedKeywords []string // Keywords that must be present
	RequiredCount    int      // Minimum number of keywords that must match
	CaseSensitive    bool     // Whether matching is case-sensitive
	AllowPartial     bool     // Allow partial word matches
}

// ValidationResult contains feedback for the player
type ValidationResult struct {
	Correct        bool          // Whether the answer is correct
	MatchedWords   []string      // Keywords that were found
	MissingWords   []string      // Keywords that were not found
	Score          float64       // 0.0 to 1.0, percentage of keywords matched
	AnnotatedInput template.HTML // User's input with matched keywords highlighted
}

// NewValidator creates a validator with default settings
func NewValidator(keywords []string, required int) *ResponseValidator {
	return &ResponseValidator{
		AcceptedKeywords: keywords,
		RequiredCount:    required,
		CaseSensitive:    false,
		AllowPartial:     false,
	}
}

// Validate checks user input against accepted keywords
func (v *ResponseValidator) Validate(input string) ValidationResult {
	if len(v.AcceptedKeywords) == 0 {
		return ValidationResult{
			Correct: true,
			Score:   1.0,
		}
	}

	normalized := normalize(input, v.CaseSensitive)

	var matched []string
	var missing []string

	for _, keyword := range v.AcceptedKeywords {
		normalizedKeyword := normalize(keyword, v.CaseSensitive)

		if v.AllowPartial {
			// Check if keyword appears anywhere as a partial match
			if strings.Contains(normalized, normalizedKeyword) {
				matched = append(matched, keyword)
			} else {
				missing = append(missing, keyword)
			}
		} else {
			// Exact word boundary matching
			if containsWord(normalized, normalizedKeyword) {
				matched = append(matched, keyword)
			} else {
				missing = append(missing, keyword)
			}
		}
	}

	score := float64(len(matched)) / float64(len(v.AcceptedKeywords))
	correct := len(matched) >= v.RequiredCount

	return ValidationResult{
		Correct:        correct,
		MatchedWords:   matched,
		MissingWords:   missing,
		Score:          score,
		AnnotatedInput: v.annotateMatches(input, matched),
	}
}

// annotateMatches highlights matched keywords in the user's input
func (v *ResponseValidator) annotateMatches(input string, matches []string) template.HTML {
	if len(matches) == 0 {
		return template.HTML(input)
	}

	// Build a regex pattern that matches any of the keywords (case-insensitive)
	// Sort by length (longest first) to avoid partial replacements
	sortedMatches := make([]string, len(matches))
	copy(sortedMatches, matches)
	sortByLength(sortedMatches)

	result := input
	for _, match := range sortedMatches {
		// Use word boundaries for exact matching
		pattern := `(?i)\b` + regexp.QuoteMeta(match) + `\b`
		re := regexp.MustCompile(pattern)

		result = re.ReplaceAllStringFunc(result, func(found string) string {
			return fmt.Sprintf(`<mark class="match-correct">%s</mark>`, found)
		})
	}

	return template.HTML(result)
}

// normalize cleans up input text for comparison
func normalize(s string, caseSensitive bool) string {
	// Remove extra whitespace
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")

	if !caseSensitive {
		s = strings.ToLower(s)
	}

	// Remove common punctuation at word boundaries
	s = strings.Trim(s, ".,!?;:")

	return s
}

// containsWord checks if a word exists with word boundaries
func containsWord(text, word string) bool {
	// Simple word boundary check - matches if word appears as a complete word
	words := strings.Fields(text)
	for _, w := range words {
		// Remove trailing punctuation
		w = strings.TrimRight(w, ".,!?;:")
		if w == word {
			return true
		}
	}
	return false
}

// sortByLength sorts strings by length (longest first)
func sortByLength(strs []string) {
	// Bubble sort is fine for small arrays
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if len(strs[i]) < len(strs[j]) {
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
}

// NumericValidator validates numeric responses
type NumericValidator struct {
	AcceptedValues []float64 // Accepted numeric values
	Tolerance      float64   // Acceptable margin of error
}

// ValidateNumeric checks if a numeric input is correct
func (v *NumericValidator) ValidateNumeric(input string) ValidationResult {
	input = strings.TrimSpace(input)

	// Try to parse as number
	var inputNum float64
	_, err := fmt.Sscanf(input, "%f", &inputNum)
	if err != nil {
		return ValidationResult{
			Correct: false,
			Score:   0.0,
		}
	}

	// Check if input matches any accepted value within tolerance
	for _, accepted := range v.AcceptedValues {
		diff := inputNum - accepted
		if diff < 0 {
			diff = -diff
		}

		if diff <= v.Tolerance {
			return ValidationResult{
				Correct:        true,
				Score:          1.0,
				AnnotatedInput: template.HTML(input),
			}
		}
	}

	return ValidationResult{
		Correct: false,
		Score:   0.0,
	}
}

// MultipleChoiceValidator validates multiple choice answers
type MultipleChoiceValidator struct {
	CorrectIndices []int // Zero-based indices of correct choices
	AllowMultiple  bool  // Whether multiple selections are allowed
}

// ValidateChoice checks if selected choices are correct
func (v *MultipleChoiceValidator) ValidateChoice(selected []int) ValidationResult {
	if !v.AllowMultiple && len(selected) > 1 {
		return ValidationResult{
			Correct: false,
			Score:   0.0,
		}
	}

	// Check if all selected choices are correct
	matchCount := 0
	for _, sel := range selected {
		correct := false
		for _, correctIdx := range v.CorrectIndices {
			if sel == correctIdx {
				correct = true
				matchCount++
				break
			}
		}
		if !correct {
			// Selected a wrong answer
			return ValidationResult{
				Correct: false,
				Score:   0.0,
			}
		}
	}

	// For multiple selection, must select ALL correct answers
	if v.AllowMultiple && matchCount != len(v.CorrectIndices) {
		return ValidationResult{
			Correct: false,
			Score:   float64(matchCount) / float64(len(v.CorrectIndices)),
		}
	}

	return ValidationResult{
		Correct: true,
		Score:   1.0,
	}
}

// Helper function to check if string is numeric
func isNumeric(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) && r != '.' && r != '-' {
			return false
		}
	}
	return true
}
