# Development Tooling

## Essential Commands

```bash
# Format code
go fmt ./...

# Check for suspicious code
go vet ./...

# Run tests
go test ./...

# Build
go build ./...

# Install
go install ./...
```

## go.mod Management

```bash
# Initialize module
go mod init github.com/user/project

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

## Static Analysis

### golangci-lint

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run
```

## Debugging

### Delve

```bash
# Install
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug
dlv debug ./main.go
```

## Code Generation

```bash
# Generate code
go generate ./...
```

## Cross-compilation

```bash
# Linux
GOOS=linux GOARCH=amd64 go build

# Windows
GOOS=windows GOARCH=amd64 go build

# macOS
GOOS=darwin GOARCH=arm64 go build
```

## Recommended Tools

- **gopls** - Go language server (IDE support)
- **goimports** - Manage imports
- **gorename** - Rename identifiers
- **guru** - Code navigation

## Next Steps

Continue to [91-testing.md](91-testing.md) to learn about testing.
