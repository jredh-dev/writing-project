# Writing Project - Technical Design

## Architecture Overview

### Technology Stack

**Backend:**
- Language: Go 1.21+
- Web Framework: net/http (standard library)
- Templating: html/template (standard library)
- Database: SQLite (for analytics)
- LLM Integration: OpenAI/Anthropic API (TBD)

**Frontend:**
- Pure HTML5
- Minimal CSS (semantic, accessible)
- Vanilla JavaScript (progressive enhancement)
- No framework dependencies

**Deployment:**
- Static assets served via CDN
- Go binary for backend
- Progressive Web App (PWA) capabilities

---

## Project Structure

```
divine-academy/
├── cmd/
│   └── prologue/          # Prologue standalone application
│       └── main.go
├── internal/
│   ├── game/              # Core game engine
│   │   ├── state.go       # Game state management
│   │   ├── choice.go      # Choice system
│   │   └── validation.go  # Answer validation
│   ├── story/             # Story content
│   │   ├── scenes.go      # Scene definitions
│   │   └── dialogue.go    # Dialogue trees
│   ├── player/            # Player data
│   │   ├── profile.go     # Player profile
│   │   ├── progress.go    # Progress tracking
│   │   └── personality.go # Personality metrics
│   └── grading/           # Grading system
│       ├── calculator.go  # Grade calculation
│       └── feedback.go    # Feedback generation
├── web/
│   ├── static/            # CSS, JS, images
│   │   ├── css/
│   │   ├── js/
│   │   └── img/
│   └── templates/         # HTML templates
│       ├── base.html
│       ├── scene.html
│       └── choice.html
├── docs/                  # Documentation
├── go.mod
├── go.sum
└── README.md
```

---

## Prologue-Specific Architecture

### Design Goals
1. **Self-Contained:** Complete story arc
2. **Frozen Codebase:** Locked after release (bug fixes only)
3. **Tutorial Foundation:** Demonstrates all core mechanics
4. **Data Collection:** Captures player responses for future use

### Core Components

#### 1. Game State Machine
```go
type GameState struct {
    CurrentScene  string
    PlayerData    *PlayerProfile
    Choices       []Choice
    Grade         *Grade
    Responses     map[string]interface{} // Dream sequence responses
}
```

#### 2. Scene System
```go
type Scene struct {
    ID          string
    Text        string
    Choices     []Choice
    ChoiceType  ChoiceType // Multiple, Many, Numeric, Open, Consequence
    Validation  ValidationFunc
    NextScene   func(*GameState) string // Dynamic branching
}
```

#### 3. Choice Types
- **MultipleChoice:** Select one option
- **ManyChoice:** Select multiple options
- **NumericResponse:** Enter number
- **OpenResponse:** Free text with validation
- **Consequence:** Show results (no player input)

#### 4. Grading System
```go
type Grade struct {
    Correct   int
    Total     int
    UsedHints bool
    Letter    string // F, C, B, A, S
}

func CalculateGrade(correct, total int, usedHints bool) Grade
```

---

## Data Flow

### Request Lifecycle

```
1. Player loads page
   ↓
2. Server renders current scene with template
   ↓
3. Player makes choice
   ↓
4. JavaScript sends POST to /choice
   ↓
5. Server validates choice
   ↓
6. Server updates game state
   ↓
7. Server calculates next scene
   ↓
8. Server renders new scene
   ↓
9. Return to step 3
```

### Save System

**Session-Based (Prologue):**
- Store game state in encrypted cookie
- No database required
- Stateless server (scales easily)
- Player can resume on same device

**Future (Post-Prologue):**
- Optional accounts
- Server-side save storage
- Cross-device sync

---

## Response Validation

### Open Response Validation

```go
type Validator struct {
    Keywords    []string      // Required words
    Patterns    []*regexp.Regexp // Regex patterns
    Fuzzy       bool          // Allow typos
    CaseInsensitive bool
}

func (v *Validator) Validate(input string) ValidationResult {
    // Extract keywords
    // Match patterns
    // Calculate similarity score
    // Return feedback
}
```

### Feedback Levels
1. **No Match:** "Try thinking about..."
2. **Partial Match:** Highlight correct keywords (CSS)
3. **Close Match:** "You're very close!"
4. **Exact Match:** Accept answer

---

## Frontend Design

### HTML Structure

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Writing Project - Prologue</title>
    <link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
    <main class="scene-container">
        <section class="narrative">
            <!-- Story text -->
        </section>
        <section class="choices">
            <!-- Choice interface -->
        </section>
    </main>
    <script src="/static/js/game.js"></script>
</body>
</html>
```

### CSS Philosophy

**Classless CSS Framework:**
- Semantic HTML styled automatically
- Minimal custom classes
- Accessibility-first
- Mobile-responsive by default

**Candidate Frameworks:**
- Water.css
- MVP.css
- Pico.css

### JavaScript Responsibilities

**Minimal JS:**
- Form submission handling
- Response validation feedback (CSS effects)
- Local progress indicators
- Optional: Typewriter effect

**Progressive Enhancement:**
- Works without JS (form posts)
- Enhanced with JS (AJAX, effects)

---

## Database Schema (Analytics Only)

### Tables

**player_sessions:**
```sql
CREATE TABLE player_sessions (
    id TEXT PRIMARY KEY,
    created_at DATETIME,
    completed_at DATETIME,
    grade TEXT,
    choices JSON,
    responses JSON
);
```

**choice_analytics:**
```sql
CREATE TABLE choice_analytics (
    scene_id TEXT,
    choice_id TEXT,
    count INTEGER,
    timestamp DATETIME
);
```

**No PII (Personally Identifiable Information) stored**

---

## Security Considerations

### Input Validation
- Sanitize all user input
- Escape HTML output
- SQL injection prevention (parameterized queries)
- XSS prevention (template escaping)

### Content Security
- HTTPS only
- CORS policies
- Rate limiting on endpoints
- Session token validation

### Privacy
- Anonymous by default
- No tracking cookies
- GDPR-compliant data handling
- Clear privacy policy

---

## Performance Targets

### Load Time
- **First Contentful Paint:** < 1 second
- **Time to Interactive:** < 2 seconds
- **Total Page Weight:** < 500KB (initial load)

### Optimization Strategies
- Gzip/Brotli compression
- Asset minification
- Image optimization (WebP with fallbacks)
- Lazy loading for non-critical assets
- Service worker for offline capability

---

## Testing Strategy

### Unit Tests
- Choice validation logic
- Grade calculation
- State transitions
- Response parsing

### Integration Tests
- Scene progression
- Save/load functionality
- Grade persistence

### Manual QA Checklist
- All choice types work
- Grades calculate correctly
- Mobile responsive
- Screen reader accessible
- Works on old browsers

---

## Deployment Strategy

### Prologue Release

**Phase 1: Local Development**
- Run on localhost
- Manual testing
- Content iteration

**Phase 2: Alpha**
- Deploy to staging server
- Invite testers
- Collect feedback

**Phase 3: Beta**
- Deploy to production
- Public access
- Analytics collection

**Phase 4: Frozen**
- Code freeze (bug fixes only)
- Documented for reference
- Foundation for future chapters

### Hosting Options
- Cloudflare Pages (static assets)
- Fly.io or Railway.app (Go backend)
- SQLite on persistent volume

---

## Future Considerations (Post-Prologue)

### Scalability
- Horizontal scaling (stateless design enables this)
- CDN for global distribution
- Database sharding if needed

### Feature Additions
- User accounts
- Chapter progression system
- Mana/combat mechanics
- Multiplayer cameos
- LLM integration

### Content Pipeline
- CMS for story content?
- Version control for narrative
- Translation workflow
- Voice acting integration

---

## Development Workflow

### Prologue Development Phases

**Phase 1: Core Engine (Week 1)**
- Game state management
- Scene system
- Basic choice types
- Template rendering

**Phase 2: Content (Week 2-3)**
- Write prologue story
- Create all scenes
- Design tutorial sequence
- Define validation rules

**Phase 3: Grading (Week 4)**
- Implement grade calculator
- Test pass/fail logic
- Feedback messages

**Phase 4: Polish (Week 5)**
- CSS styling
- JavaScript enhancements
- Mobile testing
- Accessibility audit

**Phase 5: Testing & Launch (Week 6)**
- Bug fixes
- Analytics setup
- Deploy to production
- Code freeze

---

## Code Style Guidelines

### Go Code Standards
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Comments for exported functions
- Table-driven tests
- Error handling (no panics in production)

### HTML/CSS Standards
- Semantic HTML5
- Accessible (ARIA labels where needed)
- Mobile-first responsive
- BEM naming (if custom classes needed)

### JavaScript Standards
- ES6+ syntax
- No jQuery (vanilla JS only)
- Progressive enhancement
- Minimal dependencies

---

## Monitoring & Analytics

### Metrics to Track
- Completion rate
- Average time per scene
- Choice distribution
- Grade distribution
- Drop-off points
- Browser/device breakdown

### Tools
- Custom analytics (Go backend)
- Optional: Plausible Analytics (privacy-friendly)
- Error logging (Sentry)

---

## License

**Project License:** AGPL-3.0 (per project requirements)

**Implications:**
- Source code must be open
- Derivative works must be AGPL-3.0
- Network use triggers copyleft
- Commercial use allowed but source must remain open

---

## Documentation Requirements

### Code Documentation
- README with setup instructions
- API documentation (if applicable)
- Architecture decision records (ADRs)
- Inline comments for complex logic

### Player Documentation
- How to play guide
- Controls reference
- Accessibility features
- Privacy policy
- Content warnings

