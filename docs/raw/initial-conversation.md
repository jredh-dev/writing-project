# Initial Conversation - Writing Project Project

## Date: 2025-12-27

---

## PROMPT 1: Initial Inquiry

**User:**
> how could I make a text based game with an original engine or using golang that I distribute over steam?

**Assistant Response Summary:**
- Discussed options for text-based game engines in Go
- Mentioned Steam integration requirements
- Asked clarifying questions about game type, interface, and Steam features
- Noted licensing requirements (AGPL-3.0/GPL-3.0/LGPL-3.0)

---

## PROMPT 2: Vision Clarification

**User:**
> I'd like to make a visual novel with puzzle minigames and worksheets. You're going to be going through a school, and you're going to be given assignments and questions about the world you inhabit. For instance, it'll ask you about history. There will be a point where 'hidden' options become available based on what you've learned, and you're able to type in responses. This one is for PC, but I'd like to write something like this for web/mobile. A web-based game would be cool, but I would want it to have graphics and voiceovers and be available in different languages. ideally, a website I think would be cool, because then I can do open data collection that way, and determine what routes people take. I'd like to make it as low-impact as possible so that it's accessible to old phones, bad laptops, things like that. The more I talk about it the more I just want it to be a website-based game. Each page should be a step in the story, which should only be unlockable by progress in-game. I want it to be divided into chapters, starting with a prologue that can be used as the introduction. We're going to focus on teaching the user. We want to teach them fantasy vocabulary alongside very real concepts. Most of it will be history and politics and math and things like that, which, after a certain delay, another character will provide "hints" for them. if they take the hint (it's not 100%) they'll get an 85%, which will make them a B student. At the end of the game, we reward A students for paying attention with extra content, like (i don't remember the coloquial term for secret surprises within a game) and achievements. I want that in order to pass the whole thing, you either have to have someone else play for you, or you have to learn 8th grade math/english/history/physics/chemistry/science/magic in order to get those achievements. what do you think? do you have any questions for me? I think I'm in on a pure-html / minimal CSS implementation

**Assistant Response Summary:**
- Asked 11 detailed questions about:
  - Progression & fail states
  - Hidden options mechanic
  - Hint system
  - Free-response vs multiple choice
  - Backend vs static architecture
  - Data collection scope
  - Graphics & voice
  - Localization
  - Chapter length
  - Subject integration
  - Secret surprises/unlockables
- Proposed tech stack: Go backend + HTML/CSS/Vanilla JS frontend
- Suggested phased development approach

---

## PROMPT 3: Detailed Mechanics & Requirements

**User provided comprehensive answers:**

### 1. PROGRESSION & FAIL STATES

a.) **Yes, players can fail chapters**
- Failures determined a few chapters before
- Important to bring them back to the decision that changed things
- Soft failures (don't immediate fail) should teach a lesson

b.) **Minimum passing grade is a C**
- Five grades total:
  - (F)ailure
  - (C)lass Average
  - (B)est
  - (A)ce
  - (S)uperior

c.) **Yes, players can replay chapters to improve grades**

### 2. HIDDEN OPTIONS MECHANIC

a.) **Dialogue that shows up only if you meet certain requirements**
- Less requirements = less choices
- Most important choices will have 12+ choices
- Based on factors:
  i. Individual characteristics
  ii. Attributes, aspects, and abilities (three qualities that alone should guide you to equivalent of whatever you value)
  iii. Choices made
  iv. Family around you
  v. Wealth

**Personality Tracking Examples:**
- "What is the largest number you know of?"
  - "infinity" → scientific personality
  - "999999999" → pragmatic
  - "a trillion" → practical
  - "a gazillion" → impractical
  - "five-hundred and fifty-nine thousand" → specific, calculated answer with long diatribe or answer to specific question down the line

- Food choices contribute to personality:
  - Sushi over steak
  - Steak with barbecue sauce vs steak with wine
  - Beer and burgers
  - Chicken wings
  - Salad
  - All choices contribute to personality tracking

b.) **"Learn or develop or gain enough experience with"**
- Three aspects to life: mental, emotional, physical experience
- Gaining experience in one category → get an ability point of your choice
- Ability point examples:
  - "spin move" for sports (physical)
  - "division" for mental/science
  - "keep your cool" for emotional regulation

c.) **Alternative story routes**
- Routes converge or diverge
- Don't change the path of the biggest forces (gods) JUST YET
- Plenty of room for character exploration and individual expression
- **Will make use of LLMs** - "which we've grown to hate, which I JUST WANT TO BE LIKE HEY THIS IS KINDA COOL IF USED IN THIS MANNER"
- **Multiplayer game** - you may even encounter someone's character
- If they don't "remember that" we can:
  - Insert it and gaslight them, OR
  - Insert it and let them know it was added/modified, and what this may or may not impact
- "An individual story based around real world events/timelines"

d.) **Shortcuts for people who already learned lessons**
- Example: Learning about helplessness felt during rape
  - Women and SA victims can 100% skip that
  - Should be skipped
  - To teach about the subject, involve side characters
  - They can know these details but don't have to "dive in"

### 3. HINT SYSTEM (STRINGS OF FATE)

**School Grades System:**
- Taking hints: "you're secretly a divine being, you don't need to worry"
- **Golden String of Fate**: Easiest option, most likely (with hints)
- **Red String of Fate**: Your death (by enemies)
- **White String of Fate**: Friendship (and survival)

**Real enemies should occur:**
- Example: Alliance vs Horde
- Some people hate Thrall, love Thrall, but align by him as their leader
- Story divergences like that
- People should experience different sides of the story

**LLM Customization (this phase):**
Mildly customizing the journey using LLMs to change:
1. Reading level
2. Depth
3. Intimacy
4. Style

**Post-Journey Feature:**
- Single page they can post to Instagram story
- Like a horoscope for after their journey
- Makes up "dumb lies" about what they've done today
- Example (reading when just waking up):
  > "you won two insane battles in your dreams, keeping demons from turning every single slumber into nightmares. You deserve a reward for that. Treat yourself today, and relax. Nothing is coming for you while you're awake. You've done your part keeping our world safe."

**S-Level Achievements:**
- Refusing hints is the ONLY way to activate S-level achievements

### 4. FREE-RESPONSE SYSTEM

**Keyword Highlighting Approach:**
- Let them type over and over until they find the words they want OR they give up
- Complete sentence evaluated by exact string match for ideas (nouns, places, concepts)
- When they start to tie things together: ADD CSS EFFECT the closer they get

**"Rub the Bottle" (Genie/LLM Help):**
- First three rubs are free (and necessary in tutorial)
- Each use increases timer increment:
  - 5s → 5m → 5h → 5d → 5mo → 5y
  - OR 2s → 2m → 2h → 2d → 2mo → 2y
  - OR 3s → 3m → 3h → 3d → 3mo → 3y
- Number is their "draft number" (1/100)
- s/min/h/d/mo/y is their COOLDOWN

**Reading Detection:**
- Delay is however long it takes them to read
- Include eye tracking to determine position while reading
  - Apply gentle, seizure-friendly effects
  - Color changes, droplets, etc. in iPhone version
- Natural language processing every time they ask (3 times tutorial + 5s/m/etc cooldown) to preserve token usage

**LLM Strategy:**
- Intelligently-placed use of LLM is the main achievement
- Reduce code and LLM calls as much as possible
- Opposite will be true in beginning while testing system
- Create ways for opencode to interact with story (like agentctl)
- Provide solid entertainment product for cheap

**Railings on AI Access:**
Priority order:
a.) NOT distributing illicit content to minors
b.) Obeying United States domestic security laws
c.) Anonymizing users, securing data, allowing deletion
d.) Creating a really fantastic product

**8th Grade Level Requirements:**
a.) Use no abbreviations
b.) Have proper spelling
c.) Put thought into their answer (can be relative)

**Action Mode Surprise:**
- Some people don't prioritize speaking, but action
- For them: turn story into an action game
- Type a certain word in a certain amount of time or get (metaphorically) punched in face
- If they fail ten times: record response time
- That becomes MAXIMUM SPEED for any action sequence
- Everything else will be under that

---

## DOCUMENTATION REQUEST

**User requested:**
1. Create a project
2. Create ./docs folder within project
3. Document as much detail as possible about:
   - CONCEPTS
   - GAME FEATURES
   - GAME MECHANICS
   - PLOT
   - DESIGN

**CONCEPTS:**
- Miscellaneous category
- Things that don't fall neatly into other categories
- Describe as best as possible

**GAME FEATURES include:**
- COOLDOWN mechanism
- MULTIPLE CHOICE
- OPEN RESPONSE CHOICE with LLM-LEVERAGED OPEN RESPONSE CHOICE (GENIE IN A BOTTLE)
- ABILITY POINTS after ACCUMULATING EXPERIENCE POINTS in MENTAL, PHYSICAL, or EMOTIONAL

**EXERCISES (a concept):**
Improves one of three areas:

1. **MENTAL EXERCISE**
   - Example: MATH PROBLEM (TYPE)

2. **PHYSICAL EXERCISE**
   - DDR-style or Guitar Hero-style typing (computer) or tapping (phone) exercise

3. **EMOTIONAL EXERCISE**
   - Makes you ASK QUESTIONS to relate to someone else
   - Sometimes requires LLM RESPONSE

**EMOTIONAL EXERCISE EXAMPLE:**
- Scenario: Negotiate down to FREE a SPECIAL (type) QUEST ITEM (key)
- Condition: HAVE NO MONEY (wealth condition)
- Goal: Find out what's bothering them, find creative solution
- Observation: They're twitching every time door opens with HIS BOSS checking on him

**EMPATHETIC OPTIONS:**
1. **APPEAL TO HIS HEART** (EMPATHIZE): "my boss sucks too! fuck that guy!"
2. **APPEAL TO HIS CLASS CONSCIOUSNESS** (RELATE): "fuck this place!"
3. **SEDUCE HIM** (COERCE): Bat your eyelashes if you're a PRETTY (condition) GIRL (condition)

**Three Approaches:**
- **EMPATHIZE**: Accept differences and learn from them → KARMA POSITIVE
- **RELATE**: Prove yourself to society → KARMA NEUTRAL
- **COERCE/SEDUCE**: Manipulate society → (KARMA NEGATIVE implied)

**Experience Gain:**
- Gain same amount of experience from each ENDEAVOR
- How you achieve it affects KARMA, not EXPERIENCE

---

## PROJECT NAME

**Working Title:** Writing Project (implied from "you're secretly a divine being")

---

## KEY QUOTES

> "I JUST WANT TO BE LIKE HEY THIS IS KINDA COOL IF USED IN THIS MANNER"

> "RECORD EVERYTHING THAT'S TRANSPIRED AND DOCUMENT OUR CONVERSATION AS INTRICATELY AS POSSIBLE!!! Include the raw prompts in a docs/raw directory."

---

## NEXT STEPS

1. ✅ Create project structure
2. ✅ Create docs/raw directory
3. ✅ Document raw conversation
4. [ ] Create CONCEPTS.md
5. [ ] Create GAME_FEATURES.md
6. [ ] Create GAME_MECHANICS.md
7. [ ] Create PLOT.md
8. [ ] Create DESIGN.md
