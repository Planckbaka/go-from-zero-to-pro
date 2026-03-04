# Contributing Guide

Thanks for your interest in contributing.

This repository is a learning book plus runnable examples for Go. Contributions can be:
- Documentation: new chapters, fixes, clarifications
- Examples: runnable code and tests under `examples/`
- Tooling: CI, scripts, doc build helpers
- Issue reports: typos, unclear sections, incorrect code, missing coverage

## Ground rules
- Be kind and constructive. See `CODE_OF_CONDUCT.md`.
- Prefer small PRs. One topic per PR.
- Keep things idiomatic Go. We reference:
  - Go Code Review Comments: <https://go.dev/wiki/CodeReviewComments>
  - Google Go Style Guide: <https://google.github.io/styleguide/go/>

## Repository layout
- `docs/`: the book (Markdown)
- `examples/`: runnable, testable code that supports the book
- `cmd/`: optional CLI playground and helpers
- `internal/`: internal-only packages
- `.github/workflows/`: CI

See also Go's module layout guidance: <https://go.dev/doc/modules/layout>

## How to contribute

### 1) Report an issue
Open an issue and include:
- Path and section
- Expected vs actual
- Reproduction steps (for code issues)
- Go version (`go version`) if relevant

### 2) Make a change
1. Fork the repo and create a branch:
   - `docs/<topic>`
   - `fix/<topic>`
   - `feat/<topic>`
2. Make sure code is formatted and tests pass:
   ```bash
   gofmt -w .
   go test ./...
   ```
3. If you add or modify an example:
   - Put code under `examples/<topic>/`
   - Add at least one test (`*_test.go`) or a runnable `main.go`
   - Prefer table-driven tests for clarity
   - Avoid flaky tests
4. If you add a new chapter:
   - Add `docs/<NN>-<topic>.md`
   - Include learning goals, runnable example links, common pitfalls, and optional exercises
   - Keep chapter titles and numbering consistent

### 3) Commit message conventions
Use a simple convention:
- `docs: ...`
- `examples: ...`
- `ci: ...`
- `chore: ...`
- `fix: ...`
- `feat: ...`

### 4) Pull request checklist
Before requesting review:
- [ ] `gofmt` applied
- [ ] `go test ./...` passes
- [ ] Examples updated/added are runnable
- [ ] Docs are clear and include references when appropriate
- [ ] No new lints introduced (if CI has lint step)

## Review process
Reviewers may reference Go's CodeReviewComments shorthand list:
<https://go.dev/wiki/CodeReviewComments>

Common review topics:
- Contexts and cancellation
- Error handling and error strings
- Goroutine lifetimes
- Package naming and API clarity
- Docs and examples quality

## Licensing
By contributing, you agree that your contributions will be licensed under the repository license.
