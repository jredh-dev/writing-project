# Writing Project - Quick Start Guide

## Running Your Game

### Start the Server

```bash
cd divine-academy
go run cmd/prologue/main.go
```

You'll see:
```
🎮 Writing Project Prologue running at http://localhost:8080

Open your browser and visit the URL above to play!
Press Ctrl+C to stop the server.
```

### Play the Game

1. Open http://localhost:8080 in your browser
2. You'll see your first scene with annotated text
3. Click a choice and submit
4. The story progresses to the next scene

---

## What You See (Visual Description)

### Scene View

```
┌─────────────────────────────────────────────────┐
│  Writing Project: Prologue                       │ ← Purple gradient header
└─────────────────────────────────────────────────┘

╔═══════════════════════════════════════════════╗
║  A Dream of Magic                              ║ ← Scene title (blue)
╠═══════════════════════════════════════════════╣
║                                                ║
║  You're floating in darkness. Whispers        ║
║  surround you, speaking of ancient            ║
║  [powers] and forgotten [secrets]. A voice    ║ ← Highlighted keywords
║  asks: 'What do you seek?'                    ║   (colored backgrounds)
║                                                ║
║  ┌─────────────────────────────────────────┐  ║
║  │ ○ Power to change the world            │  ║ ← Radio button choices
║  │ ○ Knowledge of the truth                │  ║   (hover effects)
║  │ ○ Peace and understanding               │  ║
║  └─────────────────────────────────────────┘  ║
║                                                ║
║          [ Continue ]                          ║ ← Submit button
║                                                ║
╚═══════════════════════════════════════════════╝
```

### Keyword Highlighting

- **Mental** (blue): `division`, `logic`, `reasoning`
- **Physical** (red): `strength`, `speed`, `endurance`  
- **Emotional** (green): `empathy`, `fear`, `joy`
- **Magic** (purple): `spell`, `wand`, `mana`

### After Submission

```
╔═══════════════════════════════════════════════╗
║  Registration Day                              ║
╠═══════════════════════════════════════════════╣
║  [Story text for next scene...]               ║
╚═══════════════════════════════════════════════╝

┌─────────────────────────────────────────────────┐
│  ✓ You chose: power                            │ ← Feedback (green box)
└─────────────────────────────────────────────────┘
```

---

## Where to Write Your Story

### Location
**File**: `internal/story/scene.go`  
**Function**: `GetPrologueScenes()`

### Example

```go
func GetPrologueScenes() []Scene {
    return []Scene{
        {
            ID:    "my-scene",
            Title: "Chapter Title",
            Text: game.NewAnnotatedText(`
                Your story text goes here.
                
                Use multiple paragraphs.
                
                Highlight important [keywords].
            `).AddKeyword("keywords", game.Mental),
            Choices: []Choice{
                {ID: "c1", Text: "Option 1", Value: "c1"},
                {ID: "c2", Text: "Option 2", Value: "c2"},
            },
        },
        // Add more scenes...
    }
}
```

---

## Story Structure Files

### Core Files

1. **`internal/story/scene.go`** ← YOU WRITE HERE
   - All your story content
   - Scene definitions
   - Dialogue
   - Choices

2. **`cmd/prologue/main.go`**
   - Scene progression logic
   - Update `getNextScene()` to define which scene follows which

3. **`web/templates/scene.html`**
   - HTML layout (rarely need to touch)

4. **`web/static/css/main.css`**
   - Styling (adjust colors, fonts if you want)

---

## Decision Tree Format

### In `scene.go`:

```go
return []Scene{
    // Scene 1
    {ID: "start", ...},
    
    // Scene 2
    {ID: "middle", ...},
    
    // Scene 3
    {ID: "end", ...},
}
```

### In `main.go`:

```go
func getNextScene(currentID, choice string) string {
    // Simple linear progression
    progression := map[string]string{
        "start":  "middle",
        "middle": "end",
    }
    
    // Or branching logic
    if currentID == "fork" {
        if choice == "left" {
            return "left-path"
        }
        return "right-path"
    }
    
    return progression[currentID]
}
```

---

## Testing Your Changes

### Rapid Iteration

1. Edit `internal/story/scene.go`
2. Save file
3. Restart server: `go run cmd/prologue/main.go`
4. Refresh browser
5. See your changes immediately

### No Database Required

Everything runs in memory. Changes are instant.

---

## Current Demo Scenes

The demo includes:

1. **dream-start** - Opening dream sequence
2. **registration** - Arrival at school
3. **campus-tour** - Meeting the campus
4. **teacher-choice** - Selecting preferences
5. **assigned-teacher** - Getting opposite choice
6. **tutorial-multiple** - Tutorial demo

Try it out, then **replace with your own story**!

---

## Features Working Now

✅ **Text rendering** with paragraph breaks  
✅ **Keyword highlighting** (4 colors)  
✅ **Multiple choice** (radio buttons)  
✅ **Scene progression** (linear or branching)  
✅ **Feedback display** after choices  
✅ **Mobile responsive** design  
✅ **Dark mode** support

## Features Coming Soon

⏳ **Open response** validation  
⏳ **Numeric input** questions  
⏳ **Many-choice** (checkboxes)  
⏳ **Grading system**  
⏳ **Player profile** tracking  
⏳ **Save/load** functionality

---

## Next Steps

1. **Try the demo**: `go run cmd/prologue/main.go`
2. **Read WRITING_GUIDE.md** for detailed story authoring
3. **Edit scene.go** to write your prologue
4. **Test and iterate** as you write

**Your passion project is ready to write!** 🎮

