# 📝 Go-Reloaded — Text Completion, Editing & Auto-Correction Tool

## 📌 Table of Contents

- [What Is This Project?](#-what-is-this-project)
- [What Does the Program Do?](#-what-does-the-program-do)
- [All the Rules Explained](#-all-the-rules-explained-with-examples)
- [Project Structure](#-project-structure)
- [How to Run the Program](#-how-to-run-the-program)
- [How to Run the Tests](#-how-to-run-the-tests)
- [How the Code Works (Step by Step)](#-how-the-code-works-step-by-step)
- [Full Examples from the Spec](#-full-examples-from-the-spec)

---

## 🎯 What Is This Project?

This is a **command-line tool** written in [Go](https://go.dev/) (also called Golang). It reads a text file, applies a series of automatic corrections and formatting rules, and writes the cleaned-up text to a new file.

Think of it like a **mini spell-checker and text formatter** — but instead of checking spelling, it follows specific rules to transform the text.

### What problem does it solve?

Imagine you have a messy text file like this:

```
I was sitting over there ,and then BAMM !!
```

The punctuation is in the wrong place — there's a space before the comma and before the exclamation marks. This program automatically fixes it to:

```
I was sitting over there, and then BAMM!!
```

It also handles number conversions, case changes, article corrections, and more!

---

## 🔧 What Does the Program Do?

The program takes **two arguments**:

1. **Input file** — the file containing the messy text
2. **Output file** — the file where the corrected text will be saved

```
go run . <input_file> <output_file>
```

That's it! The program reads the input, applies all the rules, and writes the result.

---

## 📖 All the Rules Explained (With Examples)

The program looks for special **flags** (instructions in parentheses) and other patterns in the text, and transforms the text accordingly. Here is every rule:

---

### 1. 🔢 Hexadecimal Conversion — `(hex)`

**What it does:** Finds the word `(hex)` in the text, takes the word right before it, treats that word as a [hexadecimal number](https://simple.wikipedia.org/wiki/Hexadecimal) (base 16), and converts it to a regular decimal number (base 10).

**What is hexadecimal?** It's a number system that uses 16 digits: `0-9` and `A-F`. For example, `1E` in hex = `30` in decimal.

| Before | After |
|--------|-------|
| `1E (hex) files were added` | `30 files were added` |
| `Simply add 42 (hex)` | `Simply add 66` |

---

### 2. 🔢 Binary Conversion — `(bin)`

**What it does:** Same idea as `(hex)`, but the word before is treated as a [binary number](https://simple.wikipedia.org/wiki/Binary_number) (base 2) and converted to decimal.

**What is binary?** It's a number system that uses only `0` and `1`. For example, `10` in binary = `2` in decimal.

| Before | After |
|--------|-------|
| `It has been 10 (bin) years` | `It has been 2 years` |

---

### 3. 🔠 Uppercase — `(up)` or `(up, N)`

**What it does:** Converts the word before `(up)` to ALL UPPERCASE LETTERS.

If a number is provided like `(up, 3)`, it converts the **previous 3 words** to uppercase.

| Before | After |
|--------|-------|
| `Ready, set, go (up) !` | `Ready, set, GO!` |
| `This is so exciting (up, 2)` | `This is SO EXCITING` |

---

### 4. 🔡 Lowercase — `(low)` or `(low, N)`

**What it does:** Converts the word before `(low)` to all lowercase letters.

If a number is provided like `(low, 3)`, it converts the **previous 3 words** to lowercase.

| Before | After |
|--------|-------|
| `I should stop SHOUTING (low)` | `I should stop shouting` |
| `IT WAS THE (low, 3) winter` | `it was the winter` |

---

### 5. 🔤 Capitalize — `(cap)` or `(cap, N)`

**What it does:** Converts the word before `(cap)` to have the **First Letter Capitalized** (and the rest lowercase).

If a number is provided like `(cap, 6)`, it capitalizes the **previous 6 words**.

| Before | After |
|--------|-------|
| `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| `it was the age of foolishness (cap, 6)` | `It Was The Age Of Foolishness` |

---

### 6. ✏️ Punctuation Formatting

**What it does:** Makes sure punctuation marks (`. , ! ? : ;`) are:
- **Attached** to the word before them (no space before)
- **Separated** from the word after them (space after)

Groups of punctuation like `...` or `!?` are kept together as one unit.

| Before | After |
|--------|-------|
| `I was sitting over there ,and then BAMM !!` | `I was sitting over there, and then BAMM!!` |
| `I was thinking ... You were right` | `I was thinking... You were right` |
| `what do you think ?` | `what do you think?` |

---

### 7. 💬 Quote Formatting — `' '`

**What it does:** When single quotes `'` appear as separate words, the program attaches them directly to the enclosed word(s) — no extra spaces inside the quotes.

| Before | After |
|--------|-------|
| `they describe me: ' awesome '` | `they describe me: 'awesome'` |
| `he said: ' hello world '` | `he said: 'hello world'` |

---

### 8. 📝 Article Correction — `a` → `an`

**What it does:** In English, you say "**an** apple" (not "a apple") because "apple" starts with a vowel sound. This rule automatically fixes `a` to `an` when the next word starts with a **vowel** (`a, e, i, o, u`) or the letter **`h`**.

| Before | After |
|--------|-------|
| `There it was. A amazing rock!` | `There it was. An amazing rock!` |
| `bearing a untold story` | `bearing an untold story` |

---

## 📂 Project Structure

```
go-reloaded/
├── go.mod                              # Go module file (defines the project)
├── main.go                             # Entry point — handles CLI args & file I/O
├── README.md                           # This file!
└── textprocessor/
    ├── processor.go                    # Core logic — all transformation rules
    └── processor_test.go               # Unit tests — 20 test cases
```

### What does each file do?

| File | Role | Lines | Complexity |
|------|------|-------|------------|
| `main.go` | Reads input file, calls `Process()`, writes output file | ~30 | Beginner-friendly |
| `textprocessor/processor.go` | Contains ALL the transformation rules as functions | ~220 | Intermediate |
| `textprocessor/processor_test.go` | Tests every rule with known inputs and expected outputs | ~135 | Beginner-friendly |
| `go.mod` | Tells Go this is a module named `go-reloaded` | 3 | Trivial |

---

## 🚀 How to Run the Program

### Prerequisites

You need **Go** installed on your computer. Check by running:

```bash
go version
```

If not installed, download it from [https://go.dev/dl/](https://go.dev/dl/).

### Step-by-Step

**Step 1:** Open a terminal and navigate to the project folder:

```bash
cd ~/Desktop/go-reloaded
```

**Step 2:** Create an input text file (or use any `.txt` file you have):

```bash
echo 'Simply add 42 (hex) and 10 (bin) and you will see the result is 68.' > sample.txt
```

**Step 3:** Run the program:

```bash
go run . sample.txt result.txt
```

- `go run .` — compiles and runs the Go program in the current directory
- `sample.txt` — the input file to read
- `result.txt` — the output file to create with the corrected text

**Step 4:** View the result:

```bash
cat result.txt
```

Output:

```
Simply add 66 and 2 and you will see the result is 68.
```

✅ It converted `42` (hex) to `66` and `10` (binary) to `2`!

---

## 🧪 How to Run the Tests

Tests verify that every rule works correctly. Each test provides an input string and checks that the output matches an expected result.

### Run all tests

```bash
cd ~/Desktop/go-reloaded
go test ./... -v
```

The `-v` flag means **verbose** — it shows each individual test result.

You should see output like:

```
=== RUN   TestProcess
=== RUN   TestProcess/hex_conversion
--- PASS: TestProcess/hex_conversion (0.00s)
=== RUN   TestProcess/bin_conversion
--- PASS: TestProcess/bin_conversion (0.00s)
...
PASS
ok      go-reloaded/textprocessor       0.009s
```

### Run a single test

If you want to test just one specific rule:

```bash
go test ./textprocessor -run "hex_conversion" -v
```

Replace `hex_conversion` with any test name (e.g., `punctuation_ellipsis`, `full_spec_example_1`, etc.).

---

## ⚙️ How the Code Works (Step by Step)

Here's what happens when you run `go run . sample.txt result.txt`:

### Phase 1: `main.go` — File Handling

```
1. Program starts
2. Checks that exactly 2 arguments were given (input file + output file)
3. Reads the entire input file into a string
4. Passes that string to the Process() function
5. Writes the returned result to the output file
6. Done!
```

### Phase 2: `processor.go` — The Processing Pipeline

The `Process()` function runs the text through **5 stages**, in this exact order:

```
Input Text
    │
    ▼
┌─────────────────────────────┐
│ 1. TOKENIZE                 │  Split text into individual words
│    "hello world" → ["hello",│  using whitespace as the separator
│                     "world"] │
└─────────────────────────────┘
    │
    ▼
┌─────────────────────────────┐
│ 2. APPLY FLAGS              │  Find (hex), (bin), (up), (low),
│    Scan for flag tokens and  │  (cap) and their (xx, N) variants.
│    transform preceding words │  Remove the flag tokens after
│                              │  applying the transformation.
└─────────────────────────────┘
    │
    ▼
┌─────────────────────────────┐
│ 3. FIX PUNCTUATION          │  Attach . , ! ? : ; to the
│    Move punctuation next to  │  previous word. Handle groups
│    the previous word         │  like "..." or "!?" together.
└─────────────────────────────┘
    │
    ▼
┌─────────────────────────────┐
│ 4. FIX QUOTES               │  Pair up ' marks and attach
│    ' hello ' → 'hello'      │  them to the enclosed words
│    ' a b c ' → 'a b c'      │  without extra spaces.
└─────────────────────────────┘
    │
    ▼
┌─────────────────────────────┐
│ 5. FIX ARTICLES             │  Change "a" to "an" when the
│    "a apple" → "an apple"   │  next word starts with a vowel
│    "a house" → "an house"   │  or the letter "h".
└─────────────────────────────┘
    │
    ▼
Output Text (words joined back together with spaces)
```

### Key Functions Explained

| Function | What It Does | How It Works |
|----------|-------------|--------------|
| `tokenize()` | Splits text into words | Uses `strings.Fields()` which splits by any whitespace |
| `applyFlags()` | Processes all flag commands | Uses **regex** to detect flags, then calls `applyTransformation()` |
| `applyTransformation()` | Applies one transformation | Uses a `switch` statement to pick the right action (hex→decimal, uppercase, etc.) |
| `fixPunctuation()` | Fixes punctuation spacing | **Two passes**: first splits tokens like `,what` into `,` + `what`, then attaches punctuation to the previous word |
| `fixQuotes()` | Handles quote pairing | Finds pairs of `'` tokens and wraps the words between them |
| `fixArticles()` | `a` → `an` correction | Scans for standalone `a`/`A` and checks if the next word starts with a vowel or `h` |
| `capitalize()` | Capitalizes a word | Converts to lowercase first, then uppercases the first character |

---

## 📋 Full Examples from the Spec

### Example 1 — Mixed Transformations

**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

**What happened:**
- `it (cap)` → `It` (capitalize "it")
- `times (up)` → `TIMES` (uppercase "times")
- `foolishness (cap, 6)` → capitalized the 6 words before it: `It Was The Age Of Foolishness`
- `IT WAS THE (low, 3)` → lowercased 3 words: `it was the`
- All commas were attached to the previous word

### Example 2 — Number Conversions

**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

**Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

**What happened:**
- `42 (hex)` → `42` in hexadecimal = `66` in decimal
- `10 (bin)` → `10` in binary = `2` in decimal

### Example 3 — Article Correction

**Input:**
```
There is no greater agony than bearing a untold story inside you.
```

**Output:**
```
There is no greater agony than bearing an untold story inside you.
```

**What happened:**
- `a untold` → `an untold` (because "untold" starts with a vowel)

### Example 4 — Punctuation Formatting

**Input:**
```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Output:**
```
Punctuation tests are... kinda boring, what do you think?
```

**What happened:**
- `...` was attached to "are" (group of punctuation stays together)
- `,what` — the comma was split off and attached to "boring", "what" became separate
- `?` was attached to "think"

---

## 🧰 Technologies Used

- **Language:** [Go](https://go.dev/) (version 1.21+)
- **Packages:** Only standard library packages:
  - `fmt` — formatted I/O
  - `os` — file reading/writing and CLI arguments
  - `strings` — string manipulation
  - `strconv` — string-to-number conversion
  - `regexp` — regular expressions for pattern matching
  - `unicode` — character classification (uppercase, lowercase)
  - `testing` — Go's built-in test framework

---

## 👤 Author

Built as part of the **go-reloaded** school project.
