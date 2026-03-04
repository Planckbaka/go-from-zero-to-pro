# Contributing to Go from Zero to Pro

Thank you for your interest in contributing! This document provides guidelines for contributions.

## How to Contribute

### Reporting Issues

- Check if the issue already exists
- Use a clear, descriptive title
- Include steps to reproduce (if applicable)

### Submitting Changes

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run checks: `./scripts/check.sh`
5. Commit your changes (`git commit -m 'feat: add amazing feature'`)
6. Push to your branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/go-from-zero-to-pro.git
cd go-from-zero-to-pro

# Install dependencies
go mod download

# Run checks
./scripts/check.sh
```

## Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Run `gofmt` before committing
- Add tests for new functionality
- Update documentation as needed

## Commit Messages

Use conventional commits:

- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `test:` adding/updating tests
- `refactor:` code refactoring
- `chore:` maintenance tasks

## Questions?

Open an issue with the `question` label.

Thank you for contributing! 🎉
