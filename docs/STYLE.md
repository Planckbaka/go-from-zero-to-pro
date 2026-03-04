# Project Style Guide

## Writing Style

- Use clear, concise language
- Prefer active voice
- Include runnable code examples
- Keep examples self-contained

## Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use `gofmt` for formatting
- Use `goimports` for import organization

## Documentation

- Each chapter should be a standalone learning unit
- Include "Next Steps" section linking to related chapters
- Use relative links between documents

## Examples

- Each example should be runnable
- Include tests for all examples
- Place examples in `examples/` directory matching chapter names

## Commit Messages

```
type(scope): short description

Longer description if needed.

Fixes #issue
```

Types: `docs`, `feat`, `fix`, `refactor`, `test`, `chore`
