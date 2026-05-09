# Walkthrough — Go Reloaded

This walkthrough explains the design, implementation, and testing process behind **Go Reloaded**, a text processor written in Go. It highlights the problem definition, approach, and lessons learned.

---

## 🎯 Problem Definition
The goal was to build a command-line tool that:
- Reads an input text file.
- Applies automatic corrections and transformations:
  - Hexadecimal and binary conversions.
  - Case transformations `(up)`, `(low)`, `(cap)`.
  - Punctuation normalization.
  - Quote formatting.
  - Article correction (`a` → `an`).
- Outputs the cleaned text to a new file.

---

## 🛠️ Approach
1. **Project Setup**
   - Initialized a Go module (`go.mod`).
   - Created `main.go` for CLI entry point.
   - Organized logic into `textprocessor/processor.go`.

2. **Core Logic**
   - Implemented parsing functions for each rule.
   - Used regular expressions for punctuation and article corrections.
   - Applied transformations sequentially to ensure consistency.

3. **Testing**
   - Wrote unit tests in `textprocessor/processor_test.go`.
   - Verified conversions, case changes, and punctuation handling.
   - Used `examples/sample.txt` and `examples/result.txt` for manual validation.

---

## 📂 Implementation Highlights
- **Hex/Binary Conversion**  
  Detected `(hex)` and `(bin)` markers, converted values to decimal.
- **Case Transformations**  
  Applied `(up)`, `(low)`, `(cap)` with optional counts `(up, N)`.
- **Punctuation Normalization**  
  Fixed spacing around commas, ellipsis, and question marks.
- **Quote Formatting**  
  Converted `' word '` → `'word'`.
- **Article Correction**  
  Adjusted `a` → `an` before vowels/h.

---

## 🧪 Testing Process
- **Unit Tests:**  
  Covered all transformation rules in `processor_test.go`.
- **Manual Tests:**  
  Ran the CLI with `examples/sample.txt` → `examples/result.txt`.
- **CI/CD:**  
  Added GitHub Actions workflow (`ci.yml`) to run tests automatically.

---

## 📚 Lessons Learned
- Importance of **idiomatic Go practices** for readability.
- How to structure a project with clear separation of concerns.
- Writing **unit tests** to validate edge cases.
- Setting up **CI/CD pipelines** for reliability.
- Documenting the process to make the repo recruiter‑friendly.

---

## 📌 Status
This walkthrough is part of the **Learn2Earn Fellowship** project and demonstrates my ability to design, implement, and document a Go application in a production‑style environment.
```