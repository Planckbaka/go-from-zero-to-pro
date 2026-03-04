# Security Considerations

## Input Validation

Always validate external input:

```go
import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
    return emailRegex.MatchString(email)
}
```

## SQL Injection Prevention

Use parameterized queries:

```go
// Bad - SQL injection vulnerability
query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", userInput)

// Good - parameterized
db.Query("SELECT * FROM users WHERE id = $1", userInput)
```

## XSS Prevention

Use `html/template` for HTML output:

```go
import "html/template"

tmpl := template.Must(template.New("safe").Parse(`{{.}}`))
tmpl.Execute(w, userInput) // Automatically escaped
```

## Cryptography

### Passwords

```go
import "golang.org/x/crypto/bcrypt"

// Hash
hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// Verify
err := bcrypt.CompareHashAndPassword(hash, []byte(password))
```

### Encryption

```go
import "crypto/aes"
import "crypto/cipher"
```

## Sensitive Data

### Environment Variables

```go
import "os"

dbPassword := os.Getenv("DB_PASSWORD")
```

### Don't log secrets

```go
// Bad
log.Printf("User login: %s with password %s", username, password)

// Good
log.Printf("User login: %s", username)
```

## Dependency Security

```bash
# Check for known vulnerabilities
go list -m -json all | nancy sleuth

# Use govulncheck (Go 1.18+)
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

## HTTP Security Headers

```go
func secureHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        next.ServeHTTP(w, r)
    })
}
```

## Resources

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Go Security Policy](https://go.dev/security)
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
