# GAME FEATURES - Writing Project

## Overview
This document details the concrete, player-facing features that define the Writing Project experience.

---

## GRADING SYSTEM

### Five-Tier Grade Scale

| Grade | Name | Requirement | Description |
|-------|------|-------------|-------------|
| **F** | Failure | Below 60% | Player fails chapter, must retry or return to decision point |
| **C** | Class Average | 60-74% | Minimum passing grade, standard progression |
| **B** | Best | 75-84% | Used hints (Golden String), good performance |
| **A** | Ace | 85-94% | No hints used, excellent performance |
| **S** | Superior | 95-100% | Perfect or near-perfect, unlocks secret content |

### Grade Impact:
- **C or above:** Progress to next chapter
- **B grade cap:** Maximum achievable when using hints (85%)
- **A students:** Receive extra content at game end
- **S students:** Unlock secret achievements and Easter eggs

### Replay System:
- Players can replay any chapter
- Replay to improve grade
- Grade improvements unlock additional options in future chapters
- Best grade earned is permanently recorded

---

## MULTIPLE CHOICE SYSTEM

### Standard Format:
- Present 2-12 choices depending on player conditions
- Choices affected by:
  - Individual characteristics
  - Attributes, aspects, and abilities
  - Previous choices made
  - Family relationships
  - Wealth level
  - Karma alignment

### Progressive Complexity:
- **Early Game:** 2-4 simple choices
- **Mid Game:** 4-8 nuanced choices
- **Late Game:** 8-12+ complex choices with hidden options

### Hidden Options:
- Appear only when specific requirements met
- Based on knowledge gained (e.g., learned about French Revolution → new dialogue choice)
- Based on abilities acquired
- Based on relationship levels
- Visually distinct (special color/icon?)

### Example Structure:
```
Standard visible choices:
1. Agree with the teacher
2. Question the lesson
3. Stay silent

Hidden choices (unlocked by conditions):
4. [History Knowledge] Reference the Treaty of Versailles
5. [High Charisma] Charm the class with wit
6. [Previous Choice] Mention what Sarah said yesterday
```

---

## OPEN RESPONSE SYSTEM

### Core Mechanic:
Players type complete sentences to answer questions.

### Evaluation Process:

#### 1. Keyword Extraction
- Parse input for nouns, places, concepts
- Exact string matching for key ideas
- No abbreviations allowed (8th grade standard)
- Proper spelling required

#### 2. Visual Feedback (CSS Effects)
- As player types closer to correct answer: visual effects intensify
- Effects could include:
  - Text highlighting
  - Glow effects
  - Color changes
  - Border pulses
- Seizure-friendly: gradual, smooth transitions

#### 3. Iteration
- Players can type over and over
- System highlights which parts are correct
- Can give up and take hint
- No penalty for attempts (only time/cooldown)

### Example Flow:
```
Question: "What year did World War I begin?"

Attempt 1: "sometime in the early 1900s"
→ "early 1900s" highlights green (correct range)

Attempt 2: "it started in 1914"
→ "1914" glows bright green (exact answer)
→ System accepts answer

Alternative:
Attempt 1: "idk, 1920s?"
→ No highlights, player struggles
Attempt 2: [Rubs bottle - asks Genie]
→ Genie provides hint
→ Player accepts 85% grade
```

---

## GENIE IN A BOTTLE (LLM HELP SYSTEM)

### Metaphor:
You have a magical helper (genie) you can summon for assistance.

### Tutorial Phase:
- **First 3 rubs are FREE**
- Required to complete tutorial
- Teaches mechanic safely

### Cooldown System:

#### Draft Number (1-100)
Represents your "priority" or "luck" for next request.

#### Escalating Cooldowns:
After tutorial, each rub increases wait time exponentially:

**Option A (5x):**
1. 5 seconds
2. 5 minutes
3. 5 hours
4. 5 days
5. 5 months
6. 5 years

**Option B (2x):**
1. 2 seconds
2. 2 minutes
3. 2 hours
4. 2 days
5. 2 months
6. 2 years

**Option C (3x):**
1. 3 seconds
2. 3 minutes
3. 3 hours
4. 3 days
5. 3 months
6. 3 years

### Mechanism:
- Player can ask genie specific questions about current challenge
- Natural language processing (LLM) parses question
- Provides targeted hint (not full answer)
- Caps grade at 85% (B) for that question/chapter
- Preserves token usage by limiting frequency

### Strategic Depth:
- Save hints for truly difficult questions?
- Use early and accept B grade?
- Never use and pursue S grade?

### UI Representation:
- Bottle icon or lamp icon
- Shows current cooldown if active
- Glows when available
- Shake/rub animation on activation

---

## ABILITY POINT SYSTEM

### Experience Categories:

#### Mental Experience
- Earned through: Solving puzzles, answering questions correctly, logical reasoning
- Example abilities: "Division," "Critical Thinking," "Pattern Recognition"

#### Physical Experience
- Earned through: Action sequences, rhythm challenges, quick reactions
- Example abilities: "Spin Move," "Quick Reflexes," "Endurance"

#### Emotional Experience
- Earned through: Social interactions, empathy challenges, relationship building
- Example abilities: "Keep Your Cool," "Read Emotions," "Persuasion"

### Ability Point Acquisition:
- Accumulate experience in one category
- Threshold reached → gain 1 ability point
- Choose which specific ability to unlock from available options
- Abilities unlock new options in future scenarios

### Ability Examples:

**Mental:**
- Division (solve division problems without hints)
- Chemistry Basics (identify compounds)
- Historical Context (recall dates and events)

**Physical:**
- Spin Move (dodge in action sequences)
- Fast Typing (complete sequences quicker)
- Stamina (longer action sequences)

**Emotional:**
- Keep Your Cool (resist pressure/intimidation)
- Empathy (understand others' motivations)
- Charm (improve persuasion attempts)

---

## EXERCISE SYSTEM

### Mental Exercises

**Type:** Educational mini-games embedded in curriculum

**Examples:**
- Math problems (algebra, geometry, statistics)
- Logic puzzles
- Memory challenges
- Reading comprehension
- Historical analysis

**Rewards:**
- Mental experience points
- Grade contribution
- Unlock mental abilities

### Physical Exercises

**Type:** Rhythm-based action challenges

**Mechanics:**
- **Desktop:** DDR-style or Guitar Hero-style typing
  - Words/letters appear
  - Type in rhythm
  - Accuracy and timing scored
- **Mobile:** Tapping challenges
  - Tap targets in sequence
  - Rhythm-based
  - Reaction time scored

**Adaptive Difficulty:**
- After 10 failures, system records max speed
- All future physical exercises scale to player capability
- Maintains challenge without frustration

**Rewards:**
- Physical experience points
- Unlock physical abilities
- May affect action sequences in story

### Emotional Exercises

**Type:** Social problem-solving challenges

**Mechanics:**
- Present social situation
- Player asks questions to understand NPC motivations
- Choose approach based on insights gained
- May require LLM-generated NPC responses for dynamic conversation

**Example Scenario (from initial conversation):**
```
Situation: Need to get quest item (key) from shopkeeper for free
Condition: You have no money
Observation: Shopkeeper twitches when boss enters room

Emotional Exercise:
1. Ask questions to understand shopkeeper's problem
   - "Why do you seem nervous?"
   - "Is your boss always like that?"
   - "Do you like working here?"

2. LLM generates contextual responses

3. Choose approach:
   - EMPATHIZE: "My boss sucks too! Fuck that guy!"
   - RELATE: "Fuck this place!"
   - COERCE: [If PRETTY + GIRL] Bat your eyelashes
```

**Rewards:**
- Emotional experience points
- Karma adjustment (EMPATHIZE = +karma, RELATE = neutral, COERCE = -karma)
- Same experience regardless of approach chosen
- Unlock emotional abilities

---

## CHAPTER & PROGRESSION SYSTEM

### Chapter Structure:
- **Prologue:** Introduction and tutorial
- **Numbered Chapters:** Core story progression
- **Epilogue(s):** Multiple endings based on grades and choices

### Unlocking Mechanism:
- Chapters unlock sequentially
- Must achieve C grade or higher to progress
- Can replay previous chapters anytime
- Each chapter is a discrete story beat

### Failure Recovery:
- **Soft Failures:** Teach a lesson, don't block progress
  - Example: Get poor grade on test → extra study session
- **Hard Failures:** Return to decision point
  - "Failures are usually determined a few chapters before"
  - System tracks which choice led to failure
  - Player returns to that choice point to try different path

### Save System:
- Auto-save after each choice
- Manual save at chapter breaks
- Cloud save for cross-device play (future feature)

---

## HIDDEN CONTENT & ACHIEVEMENTS

### Secret Content for A Students:
- Extra epilogue chapters
- Character backstories
- Lore codex entries
- Concept art
- Developer commentary

### S-Grade Requirements:
- Never use hints (no Golden String)
- Achieve 95%+ on all chapters
- Unlocks special achievements
- Easter eggs and secret routes

### Achievement System:
- Standard achievements for story progression
- Grade-based achievements (All A's, All S's)
- Hidden achievements (discover secrets)
- Approach-based achievements (Pure Empathy run, Pure Coercion run)
- Speed-run achievements
- Perfect memory achievements (never replay)

### Easter Eggs:
- Hidden references in text
- Secret dialogue options
- Unlockable character interactions
- Meta-commentary on game design

---

## KARMA SYSTEM

### Three Approaches to Problems:

#### EMPATHIZE (Karma Positive)
- Accept differences and learn from them
- Seek to understand others' perspectives
- Collaborative solutions
- Long-term relationship building

#### RELATE (Karma Neutral)
- Prove yourself to society
- Follow social norms
- Transactional interactions
- Practical solutions

#### COERCE/SEDUCE (Karma Negative)
- Manipulate society
- Use others for personal gain
- Short-term benefit focus
- Power-based solutions

### Key Balance:
**"Gain same amount of experience from each ENDEAVOR"**

- Karma affects story and relationships, NOT power level
- All approaches are valid gameplay
- No mechanical penalty for negative karma
- Consequences are narrative, not statistical

### Karma Tracking:
- Hidden numerical value
- Affects NPC reactions
- Unlocks karma-specific dialogue
- Influences ending variations

---

## CONDITION SYSTEM

### Player Conditions:
- **PRETTY** - Affects social interactions
- **GIRL/BOY** - Gender-specific options
- **WEALTHY/POOR** - Economic situation affects choices
- **SMART** - Unlocks intellectual options
- **STRONG** - Physical-based choices
- **CHARISMATIC** - Social persuasion options

### NPC Conditions:
- **BOSS WATCHING** - Creates specific scenarios
- **STRESSED** - Affects NPC behavior
- **FRIENDLY** - Based on relationship level
- **HOSTILE** - Based on karma/choices

### Condition Interactions:
- Multiple conditions combine
- Creates unique option sets
- "12+ choices" possible in complex scenarios
- Encourages diverse playthroughs

---

## READING DETECTION & VISUAL EFFECTS

### Eye Tracking (Mobile - Future Feature):
- Determine position in text
- Apply effects at appropriate times
- Enhanced immersion

### Visual Effects:
- **Seizure-Friendly Requirement:** All effects must be gentle
- Gradual color changes
- Droplets
- Highlight effects
- Glow effects
- Particle effects (snowflakes, leaves, embers)

### Reading Time Detection:
- Estimate reading time for text block
- Delay hint availability until player has had time to read
- Adaptive: tracks individual reading speed over time

---

## LOCALIZATION SYSTEM

### LLM-Driven Customization:

#### 1. Reading Level
- Adjust vocabulary complexity
- Simplify or elaborate sentence structure
- Maintain educational value at all levels

#### 2. Depth
- Surface-level explanations vs deep dives
- Adjust based on player performance and preferences

#### 3. Intimacy
- Casual/friendly tone vs formal tone
- Adjust based on player's interaction style

#### 4. Style
- Narrative voice adjustments
- Cultural reference adaptations
- Humor calibration

### Language Support (Future):
- Multiple language translations
- Cultural adaptation, not just translation
- Region-specific educational content where appropriate

---

## ACTION MODE (ADAPTIVE)

### Trigger:
System detects player is action-oriented rather than dialogue-oriented.

### Surprise Activation:
- Turn story into action game
- Type specific word within time limit
- Failure = metaphorical "punch in face" (health damage? Temporary stun?)

### Calibration Phase:
- Player fails action sequence 10 times
- System records their response time
- This becomes their MAXIMUM SPEED
- All future action sequences stay under this threshold

### Purpose:
- Accessibility for different play styles
- Engage players who prefer action over reading
- Dynamic difficulty adjustment
- Inclusive design (accommodate physical limitations)

---

## MULTIPLAYER INTEGRATION (ASYNCHRONOUS)

### Character Cameos:
- Encounter other players' characters as NPCs
- Their choices affect your world
- Creates sense of shared universe

### Memory Systems:

#### Gaslighting Mode:
- Insert events seamlessly
- "This always happened"
- Player discovers something "they forgot"

#### Transparent Mode:
- Notify player of insertion
- Explain what may be impacted
- Player aware it's new content

### Shared Timeline:
- Real-world events affect all players
- Individual responses diverge
- Creates discussion: "How did you handle X event?"

---

## INSTAGRAM HOROSCOPE FEATURE

### Daily Summary:
- Single shareable image/page
- Mixes truth and fantasy about player's session
- Encouraging, empowering tone

### Content Generation:
- LLM-generated based on session data
- Maintains privacy (no specific personal info)
- Shareable to social media

### Example:
> "You won two insane battles in your dreams, keeping demons from turning every single slumber into nightmares. You deserve a reward for that. Treat yourself today, and relax. Nothing is coming for you while you're awake. You've done your part keeping our world safe."

### Purpose:
- Viral marketing
- Daily engagement
- Community building
- Personalization showcase

---

## DATA COLLECTION & ANALYTICS

### Anonymous Data Collected:
- Route choices
- Grades per chapter
- Time spent
- Hint usage
- Karma alignment
- Ability choices
- Replay frequency

### Privacy Commitments:
- Anonymized collection
- User can request data deletion
- GDPR/COPPA compliant
- No personal information required
- Optional account for cross-device saves

### Usage:
- Improve game balance
- Identify confusing content
- A/B test teaching methods
- Generate insights on effective education

