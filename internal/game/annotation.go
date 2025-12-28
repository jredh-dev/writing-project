package game

import (
	"fmt"
	"html/template"
	"sort"
	"strings"
)

// KeywordCategory defines styling for different concept types
type KeywordCategory string

const (
	Mental    KeywordCategory = "mental"    // Math, logic, reasoning
	Physical  KeywordCategory = "physical"  // Action, movement, sports
	Emotional KeywordCategory = "emotional" // Feelings, empathy, relationships
	Magic     KeywordCategory = "magic"     // Fantasy-specific terms
)

// Keyword represents an annotated term with styling
type Keyword struct {
	Text     string
	Category KeywordCategory
	Learned  bool // Has player learned this concept?
}

// AnnotatedText holds text with embedded annotations
type AnnotatedText struct {
	Raw      string
	Keywords []Keyword
}

// NewAnnotatedText creates an AnnotatedText with the given raw text
func NewAnnotatedText(text string) *AnnotatedText {
	return &AnnotatedText{
		Raw:      text,
		Keywords: []Keyword{},
	}
}

// AddKeyword adds a keyword annotation to the text
func (at *AnnotatedText) AddKeyword(text string, category KeywordCategory) *AnnotatedText {
	at.Keywords = append(at.Keywords, Keyword{
		Text:     text,
		Category: category,
		Learned:  false,
	})
	return at
}

// AddLearnedKeyword adds a keyword that the player has already learned
func (at *AnnotatedText) AddLearnedKeyword(text string, category KeywordCategory) *AnnotatedText {
	at.Keywords = append(at.Keywords, Keyword{
		Text:     text,
		Category: category,
		Learned:  true,
	})
	return at
}

// Render converts annotated text to safe HTML
func (at *AnnotatedText) Render() template.HTML {
	if len(at.Keywords) == 0 {
		return template.HTML(at.Raw)
	}

	// Sort keywords by length (longest first) to avoid partial replacements
	// e.g., "long division" should be replaced before "division"
	sorted := make([]Keyword, len(at.Keywords))
	copy(sorted, at.Keywords)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i].Text) > len(sorted[j].Text)
	})

	// Build list of replacements (position-based to avoid nested replacements)
	type replacement struct {
		start int
		end   int
		html  string
	}

	replacements := []replacement{}
	lowerText := strings.ToLower(at.Raw)

	for _, kw := range sorted {
		lowerKeyword := strings.ToLower(kw.Text)
		searchStart := 0

		for {
			idx := strings.Index(lowerText[searchStart:], lowerKeyword)
			if idx == -1 {
				break
			}

			actualIdx := searchStart + idx
			endIdx := actualIdx + len(kw.Text)

			// Check if this position overlaps with existing replacements
			overlaps := false
			for _, r := range replacements {
				if (actualIdx >= r.start && actualIdx < r.end) ||
					(endIdx > r.start && endIdx <= r.end) ||
					(actualIdx <= r.start && endIdx >= r.end) {
					overlaps = true
					break
				}
			}

			if !overlaps {
				originalText := at.Raw[actualIdx:endIdx]
				marked := fmt.Sprintf(
					`<mark class="kw kw-%s%s" data-concept="%s">%s</mark>`,
					kw.Category,
					learnedClass(kw.Learned),
					kw.Text,
					originalText,
				)
				replacements = append(replacements, replacement{
					start: actualIdx,
					end:   endIdx,
					html:  marked,
				})
			}

			searchStart = actualIdx + 1
		}
	}

	// Sort replacements by start position
	sort.Slice(replacements, func(i, j int) bool {
		return replacements[i].start < replacements[j].start
	})

	// Build final result
	if len(replacements) == 0 {
		return template.HTML(at.Raw)
	}

	var result strings.Builder
	lastEnd := 0

	for _, r := range replacements {
		result.WriteString(at.Raw[lastEnd:r.start])
		result.WriteString(r.html)
		lastEnd = r.end
	}
	result.WriteString(at.Raw[lastEnd:])

	return template.HTML(result.String())
}

// learnedClass returns CSS class modifier if keyword is learned
func learnedClass(learned bool) string {
	if learned {
		return " kw-learned"
	}
	return ""
}

// RenderPlain returns the raw text without annotations (for accessibility)
func (at *AnnotatedText) RenderPlain() string {
	return at.Raw
}

// HasKeyword checks if a specific keyword is annotated
func (at *AnnotatedText) HasKeyword(text string) bool {
	lowerText := strings.ToLower(text)
	for _, kw := range at.Keywords {
		if strings.ToLower(kw.Text) == lowerText {
			return true
		}
	}
	return false
}

// MarkAsLearned marks a keyword as learned by the player
func (at *AnnotatedText) MarkAsLearned(text string) {
	lowerText := strings.ToLower(text)
	for i := range at.Keywords {
		if strings.ToLower(at.Keywords[i].Text) == lowerText {
			at.Keywords[i].Learned = true
		}
	}
}
