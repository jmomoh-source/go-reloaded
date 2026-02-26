# Project Management Assignment: Go-Reloaded Text Processing Tool

**Student:** [Your Name]  
**Course:** Project Management  
**Date:** February 26, 2026  
**Submission Deadline:** February 28, 2026

---

## 1. Project Design

### 1.1 Project Title
**Go-Reloaded** — A Text Completion, Editing & Auto-Correction Tool

### 1.2 Project Overview
Go-Reloaded is a command-line software tool built in the Go programming language. It reads a text file containing messy or improperly formatted text, applies a series of automated corrections (number conversions, case changes, punctuation fixing, quote formatting, and article correction), and writes the cleaned-up text to a new output file.

### 1.3 Problem Statement
When writing or generating text, common formatting errors frequently occur:
- Punctuation appears in the wrong place (e.g., `"hello ,world"` instead of `"hello, world"`)
- Numbers need to be converted between systems (hexadecimal/binary to decimal)
- Text casing is inconsistent
- Articles (`a`/`an`) are used incorrectly
- Quotation marks have unwanted spaces inside them

Manually correcting these errors is tedious and error-prone, especially in large documents. This project **automates** the correction process.

### 1.4 Project Objectives
| # | Objective | Success Criteria |
|---|-----------|-----------------|
| 1 | Convert hexadecimal and binary numbers to decimal | `"1E (hex)"` → `"30"` |
| 2 | Apply case transformations (uppercase, lowercase, capitalize) | `"go (up)"` → `"GO"` |
| 3 | Fix punctuation spacing | `"hello ,world"` → `"hello, world"` |
| 4 | Format single-quote pairs | `"' hello '"` → `"'hello'"` |
| 5 | Correct article usage (a/an) | `"a apple"` → `"an apple"` |
| 6 | Handle all rules in combination | Full integration test passes |

### 1.5 Target Users
- Students learning text processing
- Writers needing automated formatting cleanup
- Developers building text transformation pipelines

### 1.6 Technology Stack
| Component | Technology | Purpose |
|-----------|-----------|---------|
| Language | Go 1.21+ | Core implementation |
| Packages | Standard library only | No external dependencies |
| Testing | Go's built-in `testing` package | Quality assurance |
| Version Control | Git | Code management |

---

## 2. Principle of the Project

### 2.1 Core Principle: The Processing Pipeline
The project follows the **pipeline architecture pattern** — data flows through a series of independent transformation stages, where each stage performs one specific task and passes its output to the next stage.

```
Input File → [Tokenize] → [Apply Flags] → [Fix Punctuation] → [Fix Quotes] → [Fix Articles] → Output File
```

### 2.2 Key Design Principles

**a) Separation of Concerns**  
Each transformation rule is implemented as its own function. The hex converter doesn't know about punctuation fixing, and vice versa. This makes the code easier to understand, test, and modify.

**b) Single Responsibility**  
- `main.go` handles only file I/O (reading input, writing output)
- `processor.go` handles only text transformations
- `processor_test.go` handles only testing

**c) Order Matters**  
The 5 stages must run in a specific sequence because each stage depends on the output of the previous one:
1. **Tokenize first** — everything else needs individual words to work with
2. **Flags second** — flags like `(hex)` must be removed before punctuation fixing, otherwise they'd be treated as punctuation
3. **Punctuation third** — must happen after flags are removed to avoid misattaching punctuation
4. **Quotes fourth** — needs clean punctuation to properly pair quotes
5. **Articles last** — needs the final word list to check what comes after `"a"`

**d) Idempotency**  
Running the tool on already-correct text should produce the same text unchanged. The tool should never make correct text worse.

### 2.3 Guiding Engineering Principles
- **No external dependencies** — uses only Go's standard library for reliability and simplicity
- **Test-driven** — every rule has corresponding test cases to verify correctness
- **Fail gracefully** — invalid inputs (bad hex numbers, missing files) produce readable error messages, not crashes

---

## 3. How the Project Will Be Done

### 3.1 Development Methodology
The project follows an **incremental build approach** — each feature is built, tested, and verified before moving to the next one. This reduces risk by catching errors early.

### 3.2 Development Phases

#### Phase 1: Foundation Setup (Day 1 — 1 hour)
| Task | Deliverable |
|------|------------|
| Initialize Go module (`go mod init`) | `go.mod` |
| Create project directory structure | `textprocessor/` folder |
| Create empty `main.go` with argument validation | Working CLI skeleton |
| Create empty `processor.go` with `Process()` function | Compiles without errors |

#### Phase 2: Tokenization (Day 1 — 30 minutes)
| Task | Deliverable |
|------|------------|
| Implement `tokenize()` function | Splits text into words |
| Implement basic `Process()` that tokenizes and joins | End-to-end flow works |
| Write first test case | Test passes |

#### Phase 3: Number Conversions (Day 1 — 2 hours)
| Task | Deliverable |
|------|------------|
| Learn regex patterns for flag matching | Understanding of `regexp` package |
| Implement `applyFlags()` for `(hex)` and `(bin)` | Hex/bin conversions work |
| Implement `applyTransformation()` | Switch-based dispatch works |
| Write test cases for hex and bin | Tests pass |

#### Phase 4: Case Transformations (Day 2 — 2 hours)
| Task | Deliverable |
|------|------------|
| Add `(up)`, `(low)`, `(cap)` to `applyTransformation()` | Case changes work |
| Implement `capitalize()` helper using runes | First-letter capitalization works |
| Handle numbered flags `(up, 3)` and multi-token flags | All flag variants handled |
| Write test cases | Tests pass |

#### Phase 5: Punctuation Fixing (Day 2 — 2 hours)
| Task | Deliverable |
|------|------------|
| Implement `isPunctuation()` and `isPunctuationChar()` | Punctuation detection works |
| Implement `fixPunctuation()` two-pass algorithm | Punctuation attaches correctly |
| Write test cases for commas, ellipsis, question marks | Tests pass |

#### Phase 6: Quote Fixing (Day 2 — 1 hour)
| Task | Deliverable |
|------|------------|
| Implement `fixQuotes()` with pair-finding logic | Quotes pair correctly |
| Handle edge cases (empty quotes, unmatched quotes) | No crashes on edge cases |
| Write test cases | Tests pass |

#### Phase 7: Article Correction (Day 3 — 1 hour)
| Task | Deliverable |
|------|------------|
| Implement `isVowelOrH()` helper | Vowel detection works |
| Implement `fixArticles()` | `a` → `an` conversion works |
| Write test cases | Tests pass |

#### Phase 8: Integration & File I/O (Day 3 — 1 hour)
| Task | Deliverable |
|------|------------|
| Complete `main.go` with file reading/writing | End-to-end program works |
| Run full integration test from spec | All 19 tests pass |
| Test with various sample files | Program handles real-world input |

### 3.3 Timeline (Gantt Chart)

```
Day 1  |████ Foundation ████|██ Tokenize ██|████████ Number Conversions ████████|
Day 2  |████████ Case Transforms ████████|████████ Punctuation ████████|██ Quotes ██|
Day 3  |██ Articles ██|██ Integration ██|████ Final Testing & Review ████|
```

### 3.4 Tools & Resources Required
| Resource | Purpose |
|----------|---------|
| Go compiler (1.21+) | Building and running the project |
| Text editor / IDE (e.g., VS Code) | Writing code |
| Terminal / command line | Running and testing |
| Git | Version control and history |
| Go documentation (go.dev/doc) | Reference for standard library |

---

## 4. A to Z of the Project

A complete, exhaustive breakdown of every step from start to finish:

| Step | Activity | Details |
|------|----------|---------|
| **A** | **Analyze requirements** | Read the project specification; identify all 8 transformation rules |
| **B** | **Break down the problem** | Decompose the project into independent, testable components |
| **C** | **Create the project structure** | Initialize `go.mod`, create directories, set up file layout |
| **D** | **Design the pipeline** | Choose the 5-stage sequential processing architecture |
| **E** | **Establish the entry point** | Write `main.go` with argument parsing and file I/O |
| **F** | **First function: tokenize** | Implement `tokenize()` using `strings.Fields()` |
| **G** | **Generate test cases** | Write table-driven tests for each rule |
| **H** | **Handle hex conversion** | Implement `(hex)` flag using `strconv.ParseInt(value, 16, 64)` |
| **I** | **Implement binary conversion** | Implement `(bin)` flag using `strconv.ParseInt(value, 2, 64)` |
| **J** | **Join flags together** | Build `applyFlags()` to detect and dispatch all flag types |
| **K** | **Keen on edge cases** | Handle multi-token flags like `(cap, 6)` split across whitespace |
| **L** | **Lowercase transformation** | Implement `(low)` and `(low, N)` using `strings.ToLower()` |
| **M** | **Make uppercase work** | Implement `(up)` and `(up, N)` using `strings.ToUpper()` |
| **N** | **Nail capitalization** | Implement `(cap)` using runes for Unicode-safe first-letter uppercasing |
| **O** | **Organize punctuation logic** | Write `isPunctuation()` and `isPunctuationChar()` helpers |
| **P** | **Punctuation fixing** | Implement the two-pass `fixPunctuation()` algorithm |
| **Q** | **Quote pairing** | Implement `fixQuotes()` to find matching `'` pairs and remove inner spaces |
| **R** | **Refine article correction** | Implement `fixArticles()` for `a` → `an` before vowels and `h` |
| **S** | **Stitch the pipeline** | Wire all functions into `Process()` in the correct order |
| **T** | **Test everything** | Run `go test ./... -v` and ensure all 19 test cases pass |
| **U** | **Understand failures** | Debug any failing tests; trace input through each pipeline stage |
| **V** | **Verify with sample files** | Run the program with `sample.txt` and compare to expected `result.txt` |
| **W** | **Write documentation** | Create `README.md` explaining usage, rules, and project structure |
| **X** | **eXamine edge cases** | Test empty files, files with only punctuation, nested flags, etc. |
| **Y** | **Yesterday's review** | Review code for clarity, naming, and consistency; refactor if needed |
| **Z** | **Zero bugs confirmed** | Final full test run; project is complete and ready for submission |

---

## 5. Internal and External Factors Affecting the Project

### 5.1 Internal Factors

These are factors **within the project team's control**:

| Factor | Impact | Likelihood | Mitigation Strategy |
|--------|--------|------------|-------------------|
| **Skill level in Go** | Low proficiency slows development and increases bugs | High (for beginners) | Follow incremental learning approach; build one function at a time; use Go documentation |
| **Unclear requirements** | Misunderstanding a rule leads to wrong implementation | Medium | Re-read the spec for each rule; write tests BEFORE implementation; compare with expected outputs |
| **Poor time management** | Missing the deadline if work is backloaded | Medium | Follow the day-by-day timeline; complete at least 2 phases per day |
| **Regex complexity** | Regular expressions are hard to debug and get right | High | Test regex patterns individually using online tools (e.g., regex101.com); start simple and build up |
| **Code quality** | Messy code is harder to debug and extend | Medium | Use meaningful variable names; write comments; keep functions small and focused |
| **Testing gaps** | Missing test cases means bugs go undetected | Medium | Write tests for every rule; include edge cases (empty input, very large N, invalid hex values) |
| **Pipeline order errors** | Running stages in the wrong order produces wrong output | Low | Document and enforce the correct order in `Process()`; test the full pipeline, not just individual functions |
| **Motivation and burnout** | Working alone on a multi-day project can feel draining | Medium | Break work into small, achievable milestones; celebrate passing tests; take breaks |

### 5.2 External Factors

These are factors **outside the project team's control**:

| Factor | Impact | Likelihood | Mitigation Strategy |
|--------|--------|------------|-------------------|
| **Go version compatibility** | Code written for Go 1.21 may not work on older versions | Low | Specify `go 1.21` in `go.mod`; avoid using experimental features |
| **Operating system differences** | File paths and line endings differ between Windows, macOS, and Linux | Medium | Use `os.ReadFile()` which handles OS differences; avoid hardcoding path separators |
| **Specification ambiguity** | The project spec may not cover every edge case | Medium | Make reasonable assumptions; document them; test with the provided examples first |
| **Tool availability** | Go compiler might not be installed on the target machine | Low | Provide installation instructions in README; compile a binary for distribution |
| **Hardware limitations** | Very large input files could cause performance issues | Low | Use efficient string operations; process text as tokens rather than character-by-character |
| **Encoding issues** | Input files with non-UTF-8 encoding could cause rune conversion errors | Low | Document that UTF-8 is expected; use `[]rune` conversions for Unicode safety |
| **Peer dependency** | If working in a team, delays from one member affect everyone | Medium (if team project) | Assign clear responsibilities; use version control (Git) for parallel work |
| **Internet access** | Needed for Go documentation and regex reference tools | Low | Download Go docs offline; save key references locally |
| **Scope creep** | Temptation to add features beyond the spec (e.g., spell checking) | Medium | Stick strictly to the 8 defined rules; add extras only after core is complete and tested |
| **Deadline pressure** | External deadline of 2 days creates time pressure | High | Start immediately; follow the phased plan; prioritize core functionality over polish |

### 5.3 Risk Matrix Summary

```
            HIGH IMPACT
                │
    ┌───────────┼───────────┐
    │  Skill    │  Deadline  │
    │  Level    │  Pressure  │
    │           │            │
LOW ├───────────┼───────────┤ HIGH
LIKELIHOOD  │           │  LIKELIHOOD
    │  Encoding │  Regex     │
    │  Issues   │  Complexity│
    │           │            │
    └───────────┼───────────┘
                │
            LOW IMPACT
```

### 5.4 SWOT Analysis

| | **Helpful** | **Harmful** |
|---|-------------|------------|
| **Internal** | **Strengths:** Clear spec with examples; modular design; Go's simplicity; comprehensive tests | **Weaknesses:** Regex learning curve; pointer concepts for beginners; understanding runes vs bytes |
| **External** | **Opportunities:** Reusable tool for other text projects; Go skills transferable to industry; portfolio piece | **Threats:** Tight deadline; spec ambiguity on edge cases; OS compatibility issues |

---

## 6. Conclusion

Go-Reloaded is a well-scoped, achievable project that demonstrates core software engineering principles: modular design, pipeline architecture, test-driven development, and clean code organization. By following the phased implementation plan and managing the identified internal and external risks, the project can be completed within the 2-day deadline with high confidence.

The key success factors are:
1. **Starting immediately** — no time to waste with a 2-day deadline
2. **Building incrementally** — one function at a time, testing as you go
3. **Testing rigorously** — the 19 provided test cases are your safety net
4. **Managing complexity** — tackle simple rules first (tokenize, hex/bin) before complex ones (punctuation, quotes)

---

*Document prepared as part of the Project Management course assignment.*
