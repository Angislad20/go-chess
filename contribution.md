# ♟️ CHESS ENGINE – GO

Create a terminal-based chess engine in Go, step by step, rule by rule.

---

### Table of Contents

- [🧠 How We Work](#-how-we-work)
- [🗂️ Main Branches](#️-main-branches)
- [🌱 Branch Nomenclature](#-branch-nomenclature)
- [🛠 Contribution Steps](#-contribution-steps)

---

## 🤝 How We Work

- Tasks are created as **GitHub Issues** based on [a milestone roadmap](https://github.com/<your-github-username>/chess-go/milestones)
- Each contributor **assigns themselves to an issue**
- Each contribution **must be linked to one issue**
- All branches are created **from `dev`**
- Commit messages must **reference the issue ID**
- A **pull request (PR)** is made to `dev`
- After PR merge, the branch is **deleted**

---

## 🗂️ Main Branches

| Branch | Description |
|--------|-------------|
| `main` | Stable release (production-ready code) |
| `dev`  | Integration branch for all features |

---

## 🌱 Branch Nomenclature

```bash
Type/XX/#num-title-in-kebab-case
```

| Prefix     | When to use it                       |
|------------|--------------------------------------|
| `Feature/` | For new gameplay logic or features   |
| `Hotfix/`  | For fixing bugs or errors            |
| `Refactor/`| For code cleanup or file restructuring |

- `XX` = Your initials (e.g. `AK` for Angislad Koffi)  
- `#num` = GitHub issue number  
- `title-in-kebab-case` = brief, lowercase title separated by hyphens

### ✨ Examples

```
Feature/AK/#5-implement-pawn-movement
Hotfix/AK/#9-fix-turn-toggle-error
Refactor/AK/#12-split-move-validation
```

---

## 🛠 Contribution Steps

### 1. Fork the project

- Click on **"Fork"** at the top right of the repository's GitHub page.
- This creates a **copy of the project on your own GitHub account**.

---

### 2. Clone your fork

```bash
git clone https://github.com/<your-github-username>/chess-go.git
cd chess-go
```

---

### 3. Add remote origin to official repository

```bash
git remote add upstream https://github.com/<your-github-username>/chess-go.git
git remote -v  # to verify
git fetch upstream  # to fetch remote branches
```

---

### 4. Create a branch from issue

```bash
git checkout dev
git pull upstream dev
git checkout -b Feature/AK/#5-implement-pawn-movement
```

_Bring your contribution on your branch_

---

### 5. Commit your work

```bash
git add .
git commit -m "#5 Implement pawn movement"
```

---

### 6. Push the branch

```bash
git push -u origin Feature/AK/#5-implement-pawn-movement
```

_Proceed to a pull request (Add a clear description and link the issue with `#5`)_

---

### 7. Delete branch locally

```bash
git branch -d Feature/AK/#5-implement-pawn-movement
git push origin --delete Feature/AK/#5-implement-pawn-movement
```
