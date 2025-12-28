# Future Questions - Deferred for Post-Prologue

## Purpose
This document contains questions asked during planning that are important for future chapters but not needed for the prologue implementation.

---

## Resource System

### Mana & The Well (Q29-34)

**Q29:** What actions COST mana?
- **Partial Answer:** Spells cost mana
- **Deferred:** Specific spell costs, other mana-consuming actions

**Q30:** What is the starting mana value?
- **Answer:** 0 (until enrichment ceremony in Chapter 1)
- **Deferred:** Post-ceremony starting value

**Q31:** How much mana do different actions cost?
- **Status:** Contextual, TBD per spell/action

**Q32:** How long does it take to regenerate mana through each activity type?
- **Status:** Contextual, needs design

**Q33:** Can mana go negative (deepening The Well), or does it cap at zero?
- **Partial Answer:** The Well is always negative, mana is opposite. Spending mana when well is deep takes time to deteriorate exponentially (1hr, 2hr, 4hr...). Well never returns to 0.
- **Deferred:** Exact formula for deterioration curve

**Q34:** What happens if a player tries to do something requiring mana when they have none?
- **Answer:** They take damage. Casting without mana is dangerous (how player's parents died)
- **Deferred:** Damage amounts, death mechanics, warnings

---

## Experience & Leveling (Q35-38)

**Q35:** Are there experience point thresholds for gaining ability points?
- **Status:** TBD

**Q36:** How much XP do different activities award?
- **Status:** TBD

**Q37:** Is there a level cap, or infinite progression?
- **Status:** TBD

**Q38:** Do abilities have tiers/ranks, or are they one-time unlocks?
- **Status:** TBD

---

## Grading Algorithm (Q39-42)

**Q39:** How is a chapter grade calculated?
- **Status:** TBD (average? weighted?)

**Q40:** Do wrong answers reduce your grade, or only correct answers add to it?
- **Status:** TBD

**Q41:** Can you get partial credit?
- **Status:** TBD

**Q42:** Does time spent affect grade?
- **Status:** TBD

---

## Choice Validation (Q43-47)

**Q43:** For free-response questions, how many "correct" keyword variations should we accept per question?
- **Status:** TBD per question

**Q44:** Should capitalization matter?
- **Status:** TBD (likely no)

**Q45:** Should punctuation matter?
- **Status:** TBD (likely no for matching)

**Q46:** Do we accept misspellings within a certain edit distance?
- **Status:** TBD (accessibility vs accuracy balance)

**Q47:** What's the feedback when player is "close but not quite"?
- **Status:** TBD (CSS effects described, exact implementation TBD)

---

## Hint System Implementation (Q48-50)

**Q48:** When player "rubs the bottle," do they get multiple choice hints, or direct text hints?
- **Status:** TBD

**Q49:** Can they rub multiple times for the same question?
- **Status:** Cooldown system described, but specific behavior TBD

**Q50:** Does using a hint immediately cap that question at 85%, or the whole chapter?
- **Status:** TBD (question vs chapter scope)

---

## Personality Tracking Weights (Q51-54)

**Q51:** How many personality dimensions are we tracking total?
- **Partial Answer:** Autism score confirmed, others TBD

**Q52:** Do all choices equally impact personality, or are some weighted more?
- **Status:** TBD

**Q53:** At what point does personality "lock in" certain traits?
- **Status:** TBD

**Q54:** Can personality scores decrease, or only increase?
- **Status:** TBD

---

## Condition System Details (Q55-58)

**Q55:** How does the player learn their conditions?
- **Status:** TBD (character sheet? narrative? hidden?)

**Q56:** Can conditions change during the game?
- **Status:** TBD

**Q57:** How many conditions exist in total?
- **Status:** TBD

**Q58:** Are conditions binary or scalar?
- **Status:** TBD

---

## Karma Specifics (Q59-62)

**Q59:** What is the karma scale?
- **Status:** TBD

**Q60:** What karma value triggers different endings?
- **Status:** TBD

**Q61:** Can you see your karma score, or is it hidden?
- **Status:** TBD

**Q62:** Are there karma breakpoints that unlock specific content?
- **Status:** TBD

---

## Action Mode Calibration (Q63-65)

**Q63:** For physical exercises, what's a reasonable MAXIMUM speed?
- **Status:** TBD (WPM? TPS?)

**Q64:** What's the MINIMUM speed below which we don't go?
- **Status:** TBD (accessibility floor)

**Q65:** How many action sequences per chapter roughly?
- **Status:** TBD

---

## Educational Content (Q66-77)

**Q66:** Is each chapter focused on one subject, or mixed?
- **Partial Answer:** Prologue focused on tutorial mechanics. Future chapters: one class per chapter
- **Deferred:** Subject rotation order

**Q67:** How much content per subject across the whole game?
- **Status:** TBD

**Q68:** Are subjects taught in isolation, or integrated?
- **Status:** TBD

**Q69:** Does difficulty increase linearly, or are there spikes?
- **Status:** TBD

**Q70:** Can players choose difficulty level for specific subjects?
- **Status:** TBD

**Q71:** Are there remedial options if player consistently fails a subject?
- **Status:** TBD

**Q72:** What ratio of multiple choice vs open response vs action sequences?
- **Status:** TBD

**Q73:** Are there "exam" chapters that test multiple subjects?
- **Status:** TBD

**Q74:** Group projects or cooperative challenges?
- **Status:** TBD

**Q75:** Does the history taught match real-world 8th grade curriculum standards?
- **Status:** TBD (which country/state?)

**Q76:** Do math problems use fantasy context or real context?
- **Status:** TBD

**Q77:** Is "magic" a separate subject, or woven into physics/chemistry?
- **Status:** TBD (likely both)

---

## Design & UX (Q78-96)

**Q78:** What art style are you envisioning?
- **Status:** TBD

**Q79:** Color palette preferences?
- **Status:** TBD

**Q80:** Character portraits for NPCs - how many expressions per character?
- **Status:** TBD

**Q81:** Background art - how many unique locations?
- **Status:** TBD

**Q82:** Is this visual novel style (text at bottom, image at top)?
- **Status:** TBD

**Q83:** Where does the hint system interface appear?
- **Status:** TBD

**Q84:** How do players see their grades/progress?
- **Status:** TBD

**Q85:** Is there a "status screen" showing conditions, karma, abilities?
- **Status:** TBD

**Q86:** Should text be displayed with a "typewriter" effect, or all at once?
- **Status:** TBD

**Q87:** Font preferences?
- **Status:** TBD

**Q88:** How much text per "page"?
- **Status:** TBD

**Q89:** Voice acting for all dialogue, or just key moments?
- **Status:** TBD

**Q90:** How many languages for voice acting in Phase 1?
- **Status:** English only initially

**Q91:** Background music - adaptive based on scene mood?
- **Status:** TBD

**Q92:** Sound effects for choices, correct answers, unlocks?
- **Status:** TBD

**Q93:** Can players jump between chapters freely after unlocking?
- **Status:** TBD

**Q94:** Is there a chapter select screen?
- **Status:** TBD

**Q95:** Save slots - how many? Or auto-save only?
- **Status:** TBD

**Q96:** Can players export/share their save data?
- **Status:** TBD

---

## Technical Architecture (Q97-114)

**Q97:** Will this have user accounts, or anonymous play?
- **Status:** TBD

**Q98:** If accounts exist, what info do we collect?
- **Status:** TBD

**Q99:** How do we handle "forgot password" without email?
- **Status:** TBD

**Q100:** Where is player data stored?
- **Status:** TBD

**Q101:** Which LLM provider?
- **Status:** TBD

**Q102:** What's the acceptable latency for LLM responses?
- **Status:** TBD

**Q103:** Fallback behavior if LLM is unavailable?
- **Status:** TBD

**Q104:** How do we prevent prompt injection attacks?
- **Status:** TBD (critical security concern)

**Q105:** Real-time analytics dashboard, or batch processing?
- **Status:** TBD

**Q106:** What metrics are most important to track?
- **Status:** TBD

**Q107:** A/B testing framework needed from day one?
- **Status:** TBD

**Q108:** How do we anonymize data while still tracking individual journeys?
- **Status:** TBD

**Q109:** Max page load time target?
- **Status:** TBD

**Q110:** Max total download size for entire game?
- **Status:** TBD

**Q111:** Target frame rate for animations?
- **Status:** TBD

**Q112:** How often do player cameos appear?
- **Status:** TBD

**Q113:** Can players opt out of having their character appear in others' games?
- **Status:** TBD

**Q114:** Are cameos pre-generated, or pulled in real-time?
- **Status:** TBD

---

## Monetization & Distribution (Q115-120)

**Q115:** Is this completely free, or freemium model?
- **Status:** TBD

**Q116:** If freemium, what's paid?
- **Status:** TBD

**Q117:** Target platforms: Web only, or also mobile apps?
- **Status:** Web-based confirmed, mobile TBD

**Q118:** Distribution channels: Own domain? Steam? itch.io? App stores?
- **Status:** TBD

**Q119:** Age rating target?
- **Status:** TBD (T for Teen likely)

**Q120:** Will there be microtransactions, or one-time purchases only?
- **Status:** TBD

---

## Content Sensitivity (Q121-125)

**Q121:** Beyond SA, what other sensitive topics will appear?
- **Status:** TBD (death confirmed, others TBD)

**Q122:** How do we warn players before sensitive content?
- **Status:** TBD

**Q123:** Can players configure triggers to avoid at the start?
- **Status:** TBD

**Q124:** Should we include a content warning page before the prologue?
- **Status:** TBD

**Q125:** What's the process for adding new content warnings as we expand?
- **Status:** TBD

---

## Review Schedule

This document should be reviewed when:
1. Prologue is complete and frozen
2. Beginning Chapter 1 design
3. Designing new game systems
4. Planning major features

Questions should be moved to appropriate documentation files once answered.

