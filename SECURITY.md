# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| main    | :white_check_mark: |

## Reporting a Vulnerability

We take the security of this project seriously.

If you discover a security vulnerability, please report it by:

1. **Do not** open a public issue
2. Email the maintainer directly (or open a private security advisory on GitHub)
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

We will respond within 48 hours and provide a timeline for the fix.

## Security Best Practices

When contributing code, please follow these guidelines:

- Never commit secrets, API keys, or credentials
- Use parameterized queries for database operations
- Validate all user input
- Use `html/template` for rendering HTML
- Keep dependencies up to date
- Run `govulncheck` before submitting PRs

## Resources

- [Go Security Policy](https://go.dev/security)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
