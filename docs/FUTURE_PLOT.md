# Future Plot Points

This document captures important plot elements to be implemented in future scenes, beyond the current preface.

## Core Narrative Elements (Preserved from Original Planning)

### 1. Registration at Wincon Studiary
**Status**: Partially implemented in preface  
**Location**: `scenes/preface.yaml` - "registration" scene

The protagonist arrives at the Wincon Studiary (magic school) for their first day. This establishes:
- The urban fantasy setting (magic school in a modern city)
- The protagonist's outsider status
- The structured nature of magical education

### 2. Chance Encounter with a Student
**Status**: Partially implemented in preface  
**Location**: `scenes/preface.yaml` - "campus-tour" scene

An older, helpful student offers guidance during registration day. This:
- Introduces the social dynamics of the school
- Provides tutorial/guidance mechanics
- Could be a recurring character

### 3. Teacher Choice Paradox
**Status**: Implemented in preface  
**Location**: `scenes/preface.yaml` - "teacher-choice" and "assigned-teacher" scenes

**KEY PLOT MECHANIC**: The player is asked to choose between two teachers:
- **Professor Aldwin**: Traditional, strict, by-the-book
- **Professor Sera**: Creative, experimental, improvisational

**The Twist**: The registrar assigns the player to the OPPOSITE teacher from their choice.

**Narrative Purpose**:
- Demonstrates that player choices have unexpected consequences
- Foreshadows lack of control in the magical system
- Sets up tension between player agency and fate

### 4. The Handshake Event
**Status**: Implemented in preface  
**Location**: `scenes/preface.yaml` - "assigned-teacher" scene

During registration, the registrar shakes the protagonist's hand. Key details:
- The registrar's palm is "ice cold"
- A "strange tingling sensation" spreads up the protagonist's arm
- This causes a "sickness" (nature TBD)

**Critical Plot Point**: This handshake is what eventually turns the protagonist into the POV character for the main story.

**Questions for Future Implementation**:
- What is the nature of the "sickness"?
- How long does it take to manifest?
- Does the protagonist realize what's happening?
- Is the registrar aware they're causing this?
- What is the registrar's true identity/role?

## Implementation Notes

### Current State
The preface scenes establish:
1. Dream introduction (player motivation)
2. Registration day (setting)
3. Campus tour (world-building)
4. Teacher choice (agency vs fate theme)
5. Handshake event (transformation trigger)
6. Tutorial multiple choice (mechanics demo)

### Next Steps for Plot Development

1. **Immediate consequences of the handshake**
   - Physical symptoms?
   - Magical awareness changes?
   - Memory alterations?

2. **First day of class**
   - Meeting Professor Aldwin (the assigned teacher)
   - Introduction to classmates
   - First magic lesson

3. **Progression of the "sickness"**
   - Gradual transformation
   - Player awareness of changes
   - Other characters' reactions

4. **POV shift mechanism**
   - When does the player realize they ARE the protagonist?
   - Is there a moment of revelation?
   - How does this affect gameplay?

5. **Educational content integration**
   - Each scene should teach 8th-grade curriculum
   - Math, history, science, language arts embedded in fantasy context
   - Keywords tied to learning objectives

## Scene Writing Guidelines

When expanding beyond the preface:

1. **Use YAML format** (`scenes/*.yaml`)
2. **Annotate keywords** with appropriate categories (mental, physical, emotional, magic)
3. **Integrate curriculum** naturally into narrative
4. **Maintain mystery** around the handshake consequences
5. **Build toward POV reveal** gradually
6. **Test choices** have meaningful but not always obvious consequences

## Related Documentation

- `/scenes/preface.yaml` - Current scene definitions
- `/docs/PLOT.md` - Original plot outline
- `/docs/WRITING_GUIDE.md` - Style and tone guidelines
- `/docs/GAME_FEATURES.md` - Gameplay mechanics
