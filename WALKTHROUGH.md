# 🧠 Go-Reloaded — Complete Learning Walkthrough

This document is your coding mentor guide. It breaks down **every single concept** in the project so you can understand it deeply and then rewrite it from scratch on your own.

---

## 📌 What This Project Does (Big Picture)

This is a **text processing tool**. It reads a messy text file, applies a set of rules to clean it up, and writes the corrected text to a new file.

**Think of it as a pipeline:**
```
Messy Input Text → [5 Processing Steps] → Clean Output Text
```

The 5 steps, in order:
1. **Tokenize** — split the text into individual words
2. **Apply Flags** — handle `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` commands
3. **Fix Punctuation** — attach `. , ! ? : ;` to the previous word
4. **Fix Quotes** — pair up `'` marks and remove inner spaces
5. **Fix Articles** — change `a` to `an` before vowels/h

---

## 📂 Project Structure

```
go-reloaded/
├── go.mod                     ← Defines the Go module (like package.json in JS)
├── main.go                    ← Entry point: reads file, calls Process(), writes file
├── sample.txt                 ← Example input file
├── result.txt                 ← Example output file
└── textprocessor/
    ├── processor.go           ← ALL the logic lives here (5 functions)
    └── processor_test.go      ← Tests to verify everything works
```

---

## 🧱 File-by-File Breakdown

---

### File 1: `go.mod` (3 lines)

```go
module go-reloaded

go 1.21
```

**What you need to know:**
- Every Go project needs a `go.mod` file. It tells Go the name of your module.
- `module go-reloaded` — this is how other files import packages from this project.
- `go 1.21` — specifies the minimum Go version required.

**Go Concept: Modules**
- A **module** is a collection of Go packages. Think of it as the root of your project.
- The module name (`go-reloaded`) is used in import paths. When `main.go` imports `"go-reloaded/textprocessor"`, Go knows to look inside the `textprocessor/` folder of this module.

---

### File 2: `main.go` (32 lines) — The Entry Point

This file does **3 simple things**: validate arguments, read a file, and write a file.

#### Line-by-line:

```go
package main
```
- Every Go program starts with `package main`. The `main` package is special — it's the one that gets compiled into an executable.

```go
import (
    "fmt"
    "go-reloaded/textprocessor"
    "os"
)
```
- `fmt` — for printing messages (like `console.log` in JS)
- `os` — for file operations and command-line arguments
- `"go-reloaded/textprocessor"` — imports our custom package

**Go Concept: Imports**
- Standard library packages like `fmt` and `os` don't need a path prefix.
- Your own packages use the module name as prefix: `"go-reloaded/textprocessor"`.

```go
func main() {
```
- The `main()` function is the **entry point** — where the program starts running.

```go
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input_file> <output_file>")
        os.Exit(1)
    }
```
- `os.Args` is a slice (like an array) of command-line arguments.
- `os.Args[0]` is always the program name itself.
- So `os.Args[1]` is the input file, `os.Args[2]` is the output file.
- If the user doesn't provide exactly 2 arguments, we print usage instructions and exit with error code 1.

```go
    inputFile := os.Args[1]
    outputFile := os.Args[2]
```
- `:=` is Go's **short variable declaration**. It creates a variable and assigns a value in one step.
- Go infers the type automatically (both are `string`).

```go
    data, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
        os.Exit(1)
    }
```

**Go Concept: Error Handling**
- Go doesn't have `try/catch`. Instead, functions return **two values**: the result and an error.
- `os.ReadFile()` returns `([]byte, error)`.
- If `err != nil`, something went wrong — we print the error and exit.
- `os.Stderr` is the standard error stream (separate from normal output).
- `%v` is a format verb that prints the error's message.

```go
    result := textprocessor.Process(string(data))
```
- `string(data)` converts the byte slice to a string.
- We call our `Process()` function from the `textprocessor` package.
- **This one line is where ALL the magic happens** — the rest of `main.go` is just file I/O.

```go
    err = os.WriteFile(outputFile, []byte(result), 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
        os.Exit(1)
    }
```
- `[]byte(result)` converts the string back to bytes for writing.
- `0644` is a Unix file permission: owner can read/write, others can read.

---

### File 3: `textprocessor/processor.go` (269 lines) — The Brain

This is where ALL the text processing logic lives. Let's go function by function.

---

#### Function 1: `Process()` — The Orchestrator (Lines 12–19)

```go
func Process(text string) string {
    words := tokenize(text)
    words = applyFlags(words)
    words = fixPunctuation(words)
    words = fixQuotes(words)
    words = fixArticles(words)
    return strings.Join(words, " ")
}
```

**What it does:** Runs the text through all 5 stages in order, then joins the words back together with spaces.

**Go Concept: Exported Functions**
- In Go, functions starting with an **uppercase letter** (like `Process`) are **exported** — they can be used by other packages.
- Functions starting with a **lowercase letter** (like `tokenize`) are **private** to the package.

**Go Concept: `strings.Join()`**
- Takes a slice of strings and joins them with a separator.
- `strings.Join(["hello", "world"], " ")` → `"hello world"`

---

#### Function 2: `tokenize()` — Split Into Words (Lines 22–24)

```go
func tokenize(text string) []string {
    return strings.Fields(text)
}
```

**What it does:** Splits the text by **any whitespace** (spaces, tabs, newlines) and returns a slice of words.

**Key insight:** `strings.Fields()` is smarter than `strings.Split()` — it handles multiple spaces, tabs, and newlines automatically, and never returns empty strings.

**Example:**
```
Input:  "hello   world\tfoo"
Output: ["hello", "world", "foo"]
```

---

#### Function 3: `applyFlags()` — Process Commands (Lines 27–78)

This is the **most complex function**. It scans through all words looking for flag patterns and transforms the preceding word(s).

**Step 1: Define regex patterns**
```go
singleFlag := regexp.MustCompile(`^\((hex|bin|up|low|cap)\)$`)
numberedFlag := regexp.MustCompile(`^\((up|low|cap),\s*(\d+)\)$`)
flagStart := regexp.MustCompile(`^\((up|low|cap),$`)
flagEnd := regexp.MustCompile(`^(\d+)\)$`)
```

**Go Concept: Regular Expressions (Regex)**
- `regexp.MustCompile()` creates a compiled regex pattern. It panics if the pattern is invalid.
- The backtick `` ` `` is used for raw strings (no escape sequences needed).
- Let me explain each regex:

| Regex | Matches | Example |
|-------|---------|---------|
| `^\((hex\|bin\|up\|low\|cap)\)$` | A single flag like `(hex)` | `(hex)`, `(up)` |
| `^\((up\|low\|cap),\s*(\d+)\)$` | A numbered flag like `(up, 3)` | `(cap, 6)`, `(low,3)` |
| `^\((up\|low\|cap),$` | Start of a split flag: `(cap,` | `(cap,` |
| `^(\d+)\)$` | End of a split flag: `6)` | `6)`, `3)` |

**Why do we need 4 patterns?**
Because tokenization splits by whitespace. So `(cap, 6)` might become:
- One token: `(cap,6)` — matches `numberedFlag`
- Two tokens: `(cap,` and `6)` — matches `flagStart` + `flagEnd`

**Step 2: Loop through words**
```go
i := 0
for i < len(words) {
    word := words[i]
    // ... check patterns ...
    result = append(result, word)
    i++
}
```

**Go Concept: Manual Loop with Index**
- We use `i` instead of `range` because sometimes we need to skip tokens (when a flag spans 2 tokens, we `i += 2`).

**Step 3: When a flag is found, call `applyTransformation()`**
- Pass a **pointer** to the result slice, the flag name, and the count.
- The flag tokens themselves are NOT added to the result (they're consumed).

---

#### Function 4: `applyTransformation()` — Do the Work (Lines 81–110)

```go
func applyTransformation(result *[]string, flag string, n int) {
    r := *result
    count := n
    if count > len(r) {
        count = len(r)
    }

    startIdx := len(r) - count
    for j := startIdx; j < len(r); j++ {
        switch flag {
        case "hex":
            val, err := strconv.ParseInt(r[j], 16, 64)
            if err == nil {
                r[j] = fmt.Sprintf("%d", val)
            }
        case "bin":
            val, err := strconv.ParseInt(r[j], 2, 64)
            if err == nil {
                r[j] = fmt.Sprintf("%d", val)
            }
        case "up":
            r[j] = strings.ToUpper(r[j])
        case "low":
            r[j] = strings.ToLower(r[j])
        case "cap":
            r[j] = capitalize(r[j])
        }
    }
    *result = r
}
```

**Go Concept: Pointers**
- `result *[]string` — the `*` means this is a **pointer** to a slice. We modify the original slice, not a copy.
- `r := *result` — dereference the pointer to get the actual slice.
- `*result = r` — write the modified slice back through the pointer.

**Go Concept: `switch` Statement**
- Like `if/else if/else if` but cleaner. Go's `switch` does NOT fall through by default (no `break` needed).

**Go Concept: `strconv.ParseInt()`**
- `strconv.ParseInt("1E", 16, 64)` — parse "1E" as base-16 (hex), using 64-bit integers. Returns `30`.
- `strconv.ParseInt("10", 2, 64)` — parse "10" as base-2 (binary). Returns `2`.

**Key Logic:**
- `startIdx := len(r) - count` — start transforming from the `count`-th word from the end.
- Example: if `result` is `["This", "is", "so", "exciting"]` and `flag="up", n=2`, then `startIdx=2`, and we uppercase `"so"` and `"exciting"`.

---

#### Function 5: `capitalize()` — First Letter Upper (Lines 113–120)

```go
func capitalize(word string) string {
    if len(word) == 0 {
        return word
    }
    runes := []rune(strings.ToLower(word))
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}
```

**Go Concept: Runes**
- In Go, a `string` is a sequence of **bytes**, not characters.
- A `rune` is a single Unicode character (like `'A'`, `'é'`, `'日'`).
- `[]rune(word)` converts a string to a slice of characters, so we can safely modify individual characters.
- This matters for non-ASCII characters (e.g., `"é"` is 2 bytes but 1 rune).

---

#### Function 6: `isPunctuation()` and `isPunctuationChar()` (Lines 123–180)

```go
func isPunctuation(s string) bool {
    for _, r := range s {
        switch r {
        case '.', ',', '!', '?', ':', ';':
            continue
        default:
            return false
        }
    }
    return len(s) > 0
}
```

- Returns `true` only if **every character** in the string is a punctuation mark.
- `"..."` → `true`, `"!?"` → `true`, `"hello"` → `false`, `""` → `false`.

---

#### Function 7: `fixPunctuation()` — Two-Pass Algorithm (Lines 138–171)

```go
func fixPunctuation(words []string) []string {
```

**Pass 1: Split leading punctuation from words**
```
",what" → [",", "what"]
"!hello" → ["!", "hello"]
```

**Pass 2: Attach standalone punctuation to the previous word**
```
["there", ",", "and"] → ["there,", "and"]
["thinking", "..."] → ["thinking..."]
```

**Example walkthrough:**
```
Input words:  ["over", "there", ",and", "then", "BAMM", "!!"]
After Pass 1: ["over", "there", ",", "and", "then", "BAMM", "!!"]
After Pass 2: ["over", "there,", "and", "then", "BAMM!!"]
```

---

#### Function 8: `fixQuotes()` — Pair Up Quotes (Lines 184–225)

```go
func fixQuotes(words []string) []string {
```

**Algorithm:**
1. When we find a `'` token, search forward for the matching `'`.
2. Collect all words between the quotes.
3. Join them with spaces and wrap with `'...'`.

**Example:**
```
Input:  ["he", "said:", "'", "hello", "world", "'"]
Output: ["he", "said:", "'hello world'"]
```

**Edge cases handled:**
- Empty quotes `' '` → `''`
- No matching close quote → keeps the `'` as-is

---

#### Function 9: `fixArticles()` — a → an (Lines 228–259)

```go
func fixArticles(words []string) []string {
```

**Algorithm:**
1. Loop through all words.
2. If the current word is `"a"` or `"A"`, look at the next word.
3. Strip any leading quotes from the next word to get the actual first letter.
4. If it starts with a vowel (`a, e, i, o, u`) or `h`, change `"a"` → `"an"`.

**Go Concept: `unicode.ToLower()`**
- Converts a single rune to lowercase. Used to make the vowel check case-insensitive.

---

### File 4: `processor_test.go` (136 lines) — The Tests

```go
func TestProcess(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        // ... test cases ...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Process(tt.input)
            if got != tt.expected {
                t.Errorf(...)
            }
        })
    }
}
```

**Go Concept: Table-Driven Tests**
- This is Go's standard pattern for testing. You define a slice of test cases, each with a name, input, and expected output.
- `t.Run()` creates a **subtest** — each test case runs independently.
- `t.Errorf()` reports a failure without stopping the other tests.

**The 20 test cases cover:**

| # | What's Tested | Input → Expected |
|---|--------------|------------------|
| 1 | Hex conversion | `"1E (hex) files"` → `"30 files"` |
| 2 | Hex + Bin mixed | `"42 (hex) and 10 (bin)"` → `"66 and 2"` |
| 3 | Bin conversion | `"10 (bin) years"` → `"2 years"` |
| 4 | Uppercase single | `"go (up) !"` → `"GO!"` |
| 5 | Lowercase single | `"SHOUTING (low)"` → `"shouting"` |
| 6 | Capitalize single | `"bridge (cap)"` → `"Bridge"` |
| 7 | Uppercase N words | `"so exciting (up, 2)"` → `"SO EXCITING"` |
| 8 | Capitalize N words | `"foolishness (cap, 6)"` → 6 words capitalized |
| 9 | Lowercase N words | `"IT WAS THE (low, 3)"` → `"it was the"` |
| 10 | Punctuation comma | `",and"` → `", and"` |
| 11 | Punctuation ellipsis | `"thinking ..."` → `"thinking..."` |
| 12 | Punctuation mixed | complex punctuation test |
| 13 | Single word quotes | `"' awesome '"` → `"'awesome'"` |
| 14 | Multi word quotes | `"' hello world '"` → `"'hello world'"` |
| 15 | a→an before vowel | `"A amazing"` → `"An amazing"` |
| 16 | a→an before h | `"a huge"` → `"an huge"` |
| 17 | a stays before consonant | `"a big"` → `"a big"` |
| 18 | a→an full sentence | `"a untold"` → `"an untold"` |
| 19 | Full integration test | All rules combined |

---

## 🗺️ Rewrite Plan: Step-by-Step

Here's the order I recommend you rewrite the project, building up piece by piece:

### Phase 1: Project Setup
- [ ] Create the project folder and `go.mod`
- [ ] Create an empty `main.go` with `package main` and `func main()`
- [ ] Create the `textprocessor/` directory and an empty `processor.go`

### Phase 2: Tokenize + Join (Simplest Part)
- [ ] Write `tokenize()` using `strings.Fields()`
- [ ] Write `Process()` that only tokenizes and joins back
- [ ] Test by running: `echo "hello   world" | go run .` → `"hello world"`

### Phase 3: Number Conversions — `(hex)` and `(bin)`
- [ ] Write `applyTransformation()` with just the `hex` and `bin` cases
- [ ] Write `applyFlags()` — start with just `singleFlag` regex
- [ ] Add to `Process()` pipeline
- [ ] Write test cases for hex and bin

### Phase 4: Case Transformations — `(up)`, `(low)`, `(cap)`
- [ ] Add `up`, `low`, `cap` to `applyTransformation()`
- [ ] Write the `capitalize()` helper
- [ ] Add `numberedFlag` regex to `applyFlags()`
- [ ] Add `flagStart`/`flagEnd` for multi-token flags
- [ ] Write test cases

### Phase 5: Punctuation Fixing
- [ ] Write `isPunctuation()` and `isPunctuationChar()`
- [ ] Write `fixPunctuation()` — implement Pass 1 (split) then Pass 2 (attach)
- [ ] Add to `Process()` pipeline
- [ ] Write test cases

### Phase 6: Quote Fixing
- [ ] Write `fixQuotes()` with the pair-finding algorithm
- [ ] Add to `Process()` pipeline
- [ ] Write test cases

### Phase 7: Article Correction
- [ ] Write `isVowelOrH()`
- [ ] Write `fixArticles()`
- [ ] Add to `Process()` pipeline
- [ ] Write test cases

### Phase 8: File I/O in main.go
- [ ] Implement argument validation (`os.Args`)
- [ ] Implement file reading (`os.ReadFile`)
- [ ] Implement file writing (`os.WriteFile`)
- [ ] Test end-to-end with sample.txt

### Phase 9: Full Integration
- [ ] Run all tests: `go test ./... -v`
- [ ] Test with the full spec example from the README
- [ ] Verify edge cases work

---

## 🔑 Key Go Concepts to Master Before Starting

1. **Packages & Imports** — how Go organizes code into packages
2. **Slices** — Go's dynamic arrays (`[]string`, `append()`)
3. **Error Handling** — the `value, err := fn()` pattern
4. **Pointers** — `*` and `&` for passing by reference
5. **Range loops** — `for i, v := range slice`
6. **Runes vs Bytes** — `[]rune(str)` for Unicode-safe character operations
7. **Regular Expressions** — `regexp.MustCompile()` and `FindStringSubmatch()`
8. **Switch statements** — Go's clean alternative to `if/else` chains
9. **Testing** — `func TestXxx(t *testing.T)` and table-driven tests
10. **String manipulation** — `strings.Fields()`, `strings.Join()`, `strings.ToUpper()`

---

## 🏃 Running Commands

```bash
# Run the program
go run . sample.txt result.txt

# Run all tests (verbose)
go test ./... -v

# Run a single test
go test ./textprocessor -run "hex_conversion" -v
```
