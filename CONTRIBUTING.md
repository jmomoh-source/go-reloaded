# Contributing Guidelines

Thank you for considering contributing to **Go Reloaded**!

We welcome bug fixes, new features, and documentation improvements.  
Please follow these simple steps to ensure a smooth contribution process:

---

## 🛠️ How to Contribute

1. **Fork the repository**  
   Click the "Fork" button at the top right of this page to create your own copy.

2. **Clone your fork locally**  
   ```bash
   git clone https://github.com/<your-username>/go-reloaded.git
   cd go-reloaded
   ```

3. **Create a new branch**  
   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make your changes**  
   - Write clear, concise code.  
   - Add or update tests where necessary.  
   - Ensure your changes follow Go best practices.

5. **Run tests before committing**  
   ```bash
   go test ./... -v
   ```

6. **Commit with a descriptive message**  
   ```bash
   git commit -m "Add feature: text normalization"
   ```

7. **Push your branch**  
   ```bash
   git push origin feature/your-feature-name
   ```

8. **Open a Pull Request (PR)**  
   Go to your fork on GitHub and click "New Pull Request."  
   Provide a clear description of your changes.

---

## ✅ Code of Conduct
- Be respectful and constructive.  
- Keep discussions focused on the project.  
- Help maintain a welcoming environment for all contributors.

---

## 📌 Notes
- Major changes should be discussed in an issue before starting work.  
- Ensure CI/CD checks pass before requesting a merge.  
- Contributions are licensed under the same MIT License as the project.