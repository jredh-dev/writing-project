# Writing Project - Decision Tree & Story Structure

## How to Write Your Plot

Your story content lives in **`internal/story/scene.go`** in the `GetPrologueScenes()` function.

---

## Decision Tree Format

### Basic Scene Structure

```go
{
    ID:    "unique-scene-id",
    Title: "Scene Title (shows at top)",
    Text:  game.NewAnnotatedText("Your narrative text here..."),
    Choices: []Choice{
        {ID: "choice1", Text: "Option 1 text", Value: "choice1"},
        {ID: "choice2", Text: "Option 2 text", Value: "choice2"},
    },
}
```

### Adding Keyword Annotations

Highlight important concepts by chaining `.AddKeyword()`:

```go
Text: game.NewAnnotatedText(
    "You must learn division to understand magic theory.",
).AddKeyword("division", game.Mental).
  AddKeyword("magic theory", game.Magic),
```

**Keyword Categories:**
- `game.Mental` - Blue (math, logic, reasoning)
- `game.Physical` - Red (action, movement, sports)
- `game.Emotional` - Green (feelings, empathy, relationships)
- `game.Magic` - Purple (fantasy-specific terms)

---

## Writing Long Paragraphs

### Multi-Paragraph Text

Use string concatenation with `+` or raw strings with backticks:

```go
Text: game.NewAnnotatedText(
    "First paragraph goes here. This can be several sentences long. " +
    "You can describe the environment, the character's thoughts, whatever you need.\n\n" +
    
    "Second paragraph starts here. Use \\n\\n for paragraph breaks. " +
    "The text will render with proper spacing.\n\n" +
    
    "Third paragraph continues the story..."
),
```

**Or use raw strings (better for long text):**

```go
Text: game.NewAnnotatedText(`
    You wake to the sound of rain hammering against your window. 
    The apartment is dark, save for the flickering blue light of the 
    city outside. Today is registration day at Wincon Studiary.
    
    Your backpack sits by the door, already packed. Inside: a change 
    of clothes, the acceptance letter, and a small photo of your parents. 
    Both mages. Both gone.
    
    You grab the bag and head out into the storm.
`),
```

---

## Decision Tree Mapping

### Visual Tree Structure

I recommend mapping your story like this:

```
PROLOGUE DECISION TREE
======================

[dream-start]
├─ "Power" → personality: ambitious
├─ "Knowledge" → personality: curious  
└─ "Peace" → personality: empathetic
   ↓ (all converge)
[registration]
├─ "Excited" → mood: positive
├─ "Nervous" → mood: anxious
└─ "Cautious" → mood: guarded
   ↓ (all converge)
[campus-tour]
├─ "Accept help" → social: friendly
│  └→ [meet-student-helper] → [registration-desk]
└─ "Decline help" → social: independent
   └→ [find-it-yourself] → [registration-desk]
   
[registration-desk]
└→ [teacher-profiles]
   ↓
[teacher-choice] (OPEN RESPONSE)
├─ Prefers Aldwin → Gets Sera (opposite)
└─ Prefers Sera → Gets Aldwin (opposite)
   ↓ (consequences)
[assigned-teacher]
└→ [handshake-sickness] (critical plot point)
   ↓
[tutorial-sequence]
├→ [tutorial-multiple-choice]
├→ [tutorial-many-choice]
├→ [tutorial-numeric]
├→ [tutorial-open-response]
└→ [tutorial-consequences]
   ↓
[first-class]
... (continue story)
```

---

## Story Content Template

Here's a template for writing your scenes:

```go
// ============================================================================
// SECTION: Dream Sequence (Data Collection Phase)
// ============================================================================

{
    ID:    "dream-start",
    Title: "The Dream",
    Text: game.NewAnnotatedText(`
        [Your narrative paragraph 1 here]
        
        [Paragraph 2 with more description]
        
        [Dialogue or key moment]
    `).AddKeyword("keyword1", game.Magic).
       AddKeyword("keyword2", game.Mental),
    Choices: []Choice{
        {ID: "opt1", Text: "Choice 1 description", Value: "opt1"},
        {ID: "opt2", Text: "Choice 2 description", Value: "opt2"},
    },
},

// ============================================================================
// SECTION: Registration Day
// ============================================================================

{
    ID:    "registration",
    Title: "Arrival",
    Text: game.NewAnnotatedText(`
        [Your story continues here]
    `),
    Choices: []Choice{
        // ... choices
    },
},
```

---

## Branching Paths

### Converging Paths (Most Common)

Choices affect personality tracking but lead to same next scene:

```go
// Scene A: Make a choice
{
    ID: "choice-point",
    Choices: []Choice{
        {ID: "brave", Text: "Stand your ground", Value: "brave"},
        {ID: "flee", Text: "Run away", Value: "flee"},
    },
}

// Scene B: Both paths lead here, but text acknowledges choice
{
    ID: "after-choice",
    Text: game.NewAnnotatedText(`
        {{if eq .LastChoice "brave"}}
            You stood your ground. The creature respects your courage.
        {{else}}
            You ran, living to fight another day. Smart move.
        {{end}}
        
        Either way, you made it to safety...
    `),
}
```

### True Branching Paths (Less Common)

Different choices lead to completely different scenes:

```go
func getNextScene(currentID, choice string) string {
    switch currentID {
    case "fork-in-road":
        if choice == "left" {
            return "dark-forest"
        } else {
            return "sunny-meadow"
        }
    case "dark-forest":
        return "forest-encounter"
    case "sunny-meadow":
        return "meadow-encounter"
    default:
        return "default-next"
    }
}
```

---

## Open Response Questions

For free-text input with validation:

```go
{
    ID:    "teacher-choice",
    Title: "Choose Your Teacher",
    Text: game.NewAnnotatedText(`
        Professor Aldwin: Strict traditionalist.
        Professor Sera: Creative experimenter.
        
        Which would you prefer, and why?
    `),
    Choices: []Choice{
        {ID: "open", Text: "Type your answer...", Value: ""},
    },
    OnSubmit: func(response string) story.ValidationResult {
        validator := game.NewValidator(
            []string{"Aldwin", "Sera", "prefer", "because"}, 
            2, // Need at least 2 keywords
        )
        result := validator.Validate(response)
        
        if result.Correct {
            return story.ValidationResult{
                Valid:       true,
                Feedback:    "The registrar nods thoughtfully...",
                NextSceneID: "assigned-teacher",
            }
        }
        
        return story.ValidationResult{
            Valid:       false,
            Feedback:    "Please explain your choice in more detail.",
            NextSceneID: "teacher-choice", // Stay on same scene
        }
    },
},
```

---

## Scene Progression Logic

Update the `getNextScene()` function in `cmd/prologue/main.go`:

```go
func getNextScene(currentID, choice string) string {
    // Simple linear progression
    simpleProgression := map[string]string{
        "scene-a": "scene-b",
        "scene-b": "scene-c",
    }
    
    // Branching logic
    if currentID == "fork" {
        if choice == "left" {
            return "left-path"
        }
        return "right-path"
    }
    
    // Default linear
    if next, ok := simpleProgression[currentID]; ok {
        return next
    }
    
    return "" // End of game
}
```

---

## Tips for Writing Your Plot

### 1. **Start with the Story Beats**

List major plot points first:
- Dream (data collection)
- Registration (introduction)
- Campus tour (world building)
- Teacher choice (first consequence)
- Handshake sickness (inciting incident)
- First class (tutorial begins)
- ...

### 2. **Write Linearly First**

Don't worry about branches initially. Write the "golden path" straight through.

### 3. **Add Branches Later**

Once the main story works, add alternative paths and consequences.

### 4. **Use Comments Liberally**

```go
// ====================================
// ACT 1: ARRIVAL
// ====================================

// Scene: Player wakes up
{
    ID: "wake-up",
    // TODO: Add more sensory details
    // TODO: Reference parent's death subtly
    Text: game.NewAnnotatedText("..."),
},
```

### 5. **Track Important Choices**

Keep a list of choices that should affect later scenes:

```go
// IMPORTANT CHOICES TO TRACK:
// - dream-start: Power/Knowledge/Peace → affects hidden dialogue later
// - campus-tour: Accept/Decline help → affects social options
// - teacher-choice: Player's preference → gets opposite teacher
```

---

## Example: Complete Scene with Everything

```go
{
    ID:    "critical-moment",
    Title: "A Choice That Matters",
    Text: game.NewAnnotatedText(`
        The teacher raises her wand. Around you, other students 
        back away nervously. You remember the rule: never point 
        a wand at another person. But she's pointing it directly 
        at you.
        
        "This," she says, "is a test of your instincts."
        
        Your hand moves to your own wand before you can think. 
        The room falls silent. Everyone's watching.
        
        What do you do?
    `).AddKeyword("wand", game.Magic).
       AddKeyword("test", game.Mental).
       AddKeyword("instincts", game.Physical),
    Choices: []Choice{
        {
            ID:    "trust",
            Text:  "Trust the teacher. Keep your wand holstered.",
            Value: "trust",
        },
        {
            ID:    "defend",
            Text:  "Draw your wand in self-defense.",
            Value: "defend",
        },
        {
            ID:    "flee",
            Text:  "Duck behind your desk.",
            Value: "flee",
        },
    },
},
```

---

## Next Steps

1. **Write your scenes** in `internal/story/scene.go`
2. **Update progression** in `cmd/prologue/main.go`
3. **Test as you go**: `go run cmd/prologue/main.go`
4. **Iterate** based on how it feels to play

The system handles all the technical stuff (rendering, validation, state). 

**You just write the story!**

