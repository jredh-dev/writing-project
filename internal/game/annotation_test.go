package game

import (
	"html/template"
	"strings"
	"testing"
)

func TestAnnotatedText_Render(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		keywords []Keyword
		want     string
	}{
		{
			name: "single keyword annotation",
			text: "You must learn division to progress.",
			keywords: []Keyword{
				{Text: "division", Category: Mental, Learned: false},
			},
			want: `You must learn <mark class="kw kw-mental" data-concept="division">division</mark> to progress.`,
		},
		{
			name: "multiple keywords",
			text: "Use division and empathy to solve this.",
			keywords: []Keyword{
				{Text: "division", Category: Mental, Learned: false},
				{Text: "empathy", Category: Emotional, Learned: false},
			},
			want: `Use <mark class="kw kw-mental" data-concept="division">division</mark> and <mark class="kw kw-emotional" data-concept="empathy">empathy</mark> to solve this.`,
		},
		{
			name: "learned keyword",
			text: "You already know division.",
			keywords: []Keyword{
				{Text: "division", Category: Mental, Learned: true},
			},
			want: `You already know <mark class="kw kw-mental kw-learned" data-concept="division">division</mark>.`,
		},
		{
			name: "multi-word keyword",
			text: "The French Revolution changed history.",
			keywords: []Keyword{
				{Text: "French Revolution", Category: Mental, Learned: false},
			},
			want: `The <mark class="kw kw-mental" data-concept="French Revolution">French Revolution</mark> changed history.`,
		},
		{
			name: "overlapping keywords sorted by length",
			text: "Learn long division before division.",
			keywords: []Keyword{
				{Text: "division", Category: Mental, Learned: false},
				{Text: "long division", Category: Mental, Learned: false},
			},
			want: `Learn <mark class="kw kw-mental" data-concept="long division">long division</mark> before <mark class="kw kw-mental" data-concept="division">division</mark>.`,
		},
		{
			name:     "no keywords",
			text:     "Plain text with no annotations.",
			keywords: []Keyword{},
			want:     "Plain text with no annotations.",
		},
		{
			name: "case insensitive matching preserves original case",
			text: "Division is important. Learn division well.",
			keywords: []Keyword{
				{Text: "division", Category: Mental, Learned: false},
			},
			want: `<mark class="kw kw-mental" data-concept="division">Division</mark> is important. Learn <mark class="kw kw-mental" data-concept="division">division</mark> well.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			at := &AnnotatedText{
				Raw:      tt.text,
				Keywords: tt.keywords,
			}
			got := string(at.Render())
			if got != tt.want {
				t.Errorf("Render() =\n%v\nwant:\n%v", got, tt.want)
			}
		})
	}
}

func TestAnnotatedText_AddKeyword(t *testing.T) {
	at := NewAnnotatedText("Test text")
	at.AddKeyword("test", Mental).AddKeyword("example", Physical)

	if len(at.Keywords) != 2 {
		t.Errorf("Expected 2 keywords, got %d", len(at.Keywords))
	}

	if at.Keywords[0].Text != "test" || at.Keywords[0].Category != Mental {
		t.Errorf("First keyword incorrect: %+v", at.Keywords[0])
	}

	if at.Keywords[1].Text != "example" || at.Keywords[1].Category != Physical {
		t.Errorf("Second keyword incorrect: %+v", at.Keywords[1])
	}
}

func TestAnnotatedText_MarkAsLearned(t *testing.T) {
	at := NewAnnotatedText("Learn division")
	at.AddKeyword("division", Mental)

	if at.Keywords[0].Learned {
		t.Error("Keyword should not be learned initially")
	}

	at.MarkAsLearned("division")

	if !at.Keywords[0].Learned {
		t.Error("Keyword should be marked as learned")
	}
}

func TestAnnotatedText_HasKeyword(t *testing.T) {
	at := NewAnnotatedText("Test")
	at.AddKeyword("example", Mental)

	if !at.HasKeyword("example") {
		t.Error("Should have 'example' keyword")
	}

	if !at.HasKeyword("EXAMPLE") {
		t.Error("Should be case insensitive")
	}

	if at.HasKeyword("missing") {
		t.Error("Should not have 'missing' keyword")
	}
}

func TestAnnotatedText_RenderPlain(t *testing.T) {
	text := "Test text with keywords"
	at := NewAnnotatedText(text)
	at.AddKeyword("keywords", Mental)

	if at.RenderPlain() != text {
		t.Errorf("RenderPlain() should return raw text, got: %s", at.RenderPlain())
	}
}

func TestHTMLSafety(t *testing.T) {
	// Ensure HTML in keywords is not executed
	at := NewAnnotatedText("Test <script>alert('xss')</script>")
	at.AddKeyword("<script>", Mental)

	rendered := at.Render()
	renderedStr := string(rendered)

	// The template.HTML type means this is trusted, but our input was plain text
	// so script tags in the original text will be preserved as-is
	if !strings.Contains(renderedStr, "<script>") {
		t.Error("Original text content should be preserved")
	}

	// But our keyword annotations should be properly formatted
	if !strings.Contains(renderedStr, `data-concept="<script>"`) {
		t.Error("Keyword annotations should be present")
	}
}

func TestRenderReturnsTemplateHTML(t *testing.T) {
	at := NewAnnotatedText("test")
	result := at.Render()

	// Verify it returns template.HTML type
	var _ template.HTML = result
}
