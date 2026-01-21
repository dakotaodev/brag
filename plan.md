# brag-cli Implementation Backlog

This document tracks the granular tasks required to build a resume-ready CLI tool in Go. Each task is designed to teach a specific Go concept.

---

## 🏗 Phase 1: Foundations (The MVP)
*Goal: Master Go types and basic file interaction.*

- [ ] **Task 1.1: Project Initialization**
    - Run `go mod init github.com/[USER]/brag-cli`.
    - Create a basic `main.go` that prints "Brag CLI Initialized".
    - **Learning:** Go Modules and workspace setup.
- [ ] **Task 1.2: Define the Brag Schema**
    - Create an `internal/models` package.
    - Define a `Brag` struct with JSON tags (e.g., ``json:"content"``).
    - **Learning:** Structs, Types, and Struct Tags.
- [ ] **Task 1.3: JSON Storage Engine**
    - Write a function `SaveBrag(b Brag)` that reads `brags.json`, appends the new struct, and writes it back.
    - Handle file creation if it doesn't exist.
    - **Learning:** `os` and `encoding/json` packages, Error handling (`if err != nil`).
- [ ] **Task 1.4: Primitive CLI Flags**
    - Use the `flag` package to accept a `-msg` string from the terminal.
    - **Learning:** Capturing CLI arguments.

---

## 🏗 Phase 2: Professional Architecture
*Goal: Implement industry-standard CLI patterns.*

- [ ] **Task 2.1: Bootstrap Cobra Framework**
    - Install Cobra: `go get github.com/spf13/cobra`.
    - Initialize Cobra structure: `cobra-cli init`.
    - **Learning:** Third-party package management.
- [ ] **Task 2.2: Implement `brag add` Command**
    - Move your Phase 1 logic into the `add` subcommand.
    - Add flags for `--category` and `--impact`.
    - **Learning:** Cobra command routing and flag persistence.
- [ ] **Task 2.3: Implement `brag list` Command**
    - Create logic to iterate through the JSON file and print entries.
    - Format output using the `text/tabwriter` package for clean columns.
    - **Learning:** Slices, Range loops, and Standard Library formatting.
- [ ] **Task 2.4: Global Config Path**
    - Use `os.UserHomeDir()` to ensure the data file lives in `~/.brag-cli/data.json` rather than the local folder.
    - **Learning:** Environment variables and cross-platform file paths.

---

## 🏗 Phase 3: Enhanced Features & UX
*Goal: Build features that make the tool "sticky" and user-friendly.*

- [ ] **Task 3.1: Interactive Prompting**
    - Use `github.com/AlecAivazis/survey/v2` to create a fallback prompt if no flags are provided during `brag add`.
    - **Learning:** Handling interactive I/O.
- [ ] **Task 3.2: Markdown Template Engine**
    - Create a `templates/` folder with a `.md` template.
    - Use `text/template` to map your JSON data into a formatted "Performance Review" document.
    - **Learning:** Go's powerful templating engine.
- [ ] **Task 3.3: Data Filtering Logic**
    - Add a `--days` flag to