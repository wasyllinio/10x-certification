# Development Guide

## Quick Start

### 1. Install Development Tools

```bash
cd backend
bash .golangci-install.sh
```

### 2. Format and Check Your Code

```bash
# Format code
make fmt

# Check for issues
make lint

# Auto-fix issues
make lint-fix

# Run all checks
make check
```

### 3. Run Tests

```bash
# Run tests
make test

# With coverage report
make test-coverage
```

## Available Commands

Run `make help` to see all available commands.

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |
| `make lint` | Run golangci-lint analysis |
| `make lint-fix` | Run golangci-lint with auto-fix |
| `make fmt` | Format code with gofmt and goimports |
| `make vet` | Run go vet |
| `make test` | Run tests |
| `make test-coverage` | Run tests with coverage report |
| `make check` | Run all checks (lint + vet + test) |
| `make security` | Run security scan with gosec |
| `make install-tools` | Install development tools |
| `make build` | Build the application |
| `make run` | Run the application |
| `make clean` | Clean build artifacts |
| `make all` | Run everything (fmt + lint + vet + test) |
| `make ci` | Run CI checks (lint + vet + test) |

## Code Quality Standards

### Linting

The project uses **golangci-lint** which aggregates 50+ linters. Configuration is in `.golangci.yml`.

Key linters enabled:
- **Critical**: errcheck, staticcheck, govet, typecheck
- **Code Quality**: gocritic, goconst, misspell, unconvert
- **Complexity**: gocyclo, dupl, funlen
- **Security**: gosec

### Before Committing

Always run checks before committing:

```bash
make check
```

To format and auto-fix issues:

```bash
make all
```

## VS Code Integration

The workspace is pre-configured with:

- ✅ Auto-formatting on save
- ✅ Import organization on save
- ✅ Linting with golangci-lint
- ✅ Vet check on save
- ✅ Linting only modified files

### Required Extensions

Install these VS Code extensions for the best experience:

1. [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
2. [golangci-lint extension](https://marketplace.visualstudio.com/items?itemName=Golang.go)

## CI/CD

GitHub Actions automatically runs:

- ✅ Linting with golangci-lint
- ✅ Static analysis with go vet
- ✅ Unit tests with coverage report
- ✅ Security scan with gosec
- ✅ Build verification

These checks run on every PR and push to `main`/`develop`.

## Troubleshooting

### golangci-lint not found

```bash
make install-tools
```

### VS Code not showing errors

1. Install "golangci-lint for VS Code" extension
2. Restart VS Code
3. Ensure golangci-lint is in PATH: `which golangci-lint`

### Make not found

```bash
# macOS
brew install make

# Linux
sudo apt-get install build-essential
```

## Best Practices

1. **Always run `make check` before committing**
2. **Use `make lint-fix` to auto-fix issues when possible**
3. **Write meaningful commit messages**
4. **Keep functions simple** (use gocyclo warnings as a guide)
5. **Handle all errors** (errcheck is critical)
6. **Run tests before pushing** (`make test`)

## Project Structure

This project follows Domain-Driven Design (DDD) and CQRS patterns:

```
backend/
├── cmd/api/           # Application entry point
├── internal/
│   ├── application/   # Application layer (container, DI)
│   ├── domain/        # Domain layer (business logic)
│   │   ├── auth/      # Auth domain
│   │   ├── chargers/  # Chargers domain
│   │   └── locations/ # Locations domain
│   ├── infrastructure/# Infrastructure layer
│   └── shared/        # Shared utilities
└── migrations/        # Database migrations
```

Each domain contains:
- `command/` - Write operations
- `query/` - Read operations
- `model/` - Domain models
- `repository/` - Data access interfaces
- `service/` - Domain services
- `dto/` - Data Transfer Objects
- `events/` - Domain events

