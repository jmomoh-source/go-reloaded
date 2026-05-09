# Project Management — Go Reloaded

This document outlines the planning, risk analysis, and reflection process for the **Go Reloaded** project. It demonstrates how project management principles were applied to a technical build.

---

## 🎯 Objectives
- Build a Go CLI tool for text processing and correction.
- Apply structured project management practices:
  - Requirement gathering
  - Task breakdown
  - Risk analysis
  - Iterative development
  - Reflection

---

## 📌 Requirements
- Input: text file (`sample.txt`)
- Output: cleaned text file (`result.txt`)
- Rules implemented:
  - Hexadecimal and binary conversions
  - Case transformations `(up)`, `(low)`, `(cap)`
  - Punctuation normalization
  - Quote formatting
  - Article correction

---

## 🗂️ Task Breakdown
1. **Setup**
   - Initialize Go module (`go.mod`)
   - Create project structure
2. **Core Development**
   - Implement text transformations in `processor.go`
   - Add CLI entry point (`main.go`)
3. **Testing**
   - Unit tests in `processor_test.go`
   - Manual validation with `examples/`
4. **Documentation**
   - `README.md` for repo overview
   - `WALKTHROUGH.md` for technical process
   - `PROJECT_MANAGEMENT.md` for planning
5. **CI/CD**
   - GitHub Actions workflow (`ci.yml`) for automated testing

---

## ⚠️ Risk Analysis
- **Complexity of rules** → Mitigated by modular functions.
- **Edge cases in text processing** → Addressed with unit tests.
- **Time constraints** → Managed by breaking tasks into small milestones.
- **Tooling unfamiliarity (Go)** → Reduced by consulting official docs and Effective Go.

---

## 🔄 Iterative Development
- Started with simple transformations (case changes).
- Gradually added complexity (hex/bin conversions, punctuation).
- Tested each feature before moving to the next.
- Maintained documentation throughout.

---

## 📚 Reflection
- Learned the importance of **clear requirements** before coding.
- Reinforced the value of **unit testing** for reliability.
- Gained experience in **Go project structure** and idiomatic practices.
- Understood how **CI/CD pipelines** improve confidence in code quality.
- Practiced documenting both technical and management aspects for recruiter visibility.

---

## 📌 Status
This project was completed as part of the **Learn2Earn Fellowship**.  
It demonstrates not only technical ability in Go but also the application of project management principles to software development.
```