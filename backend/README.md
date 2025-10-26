# Backend - EV Charger Management System

## Konfiguracja

Aplikacja używa [Viper](https://github.com/spf13/viper) do zarządzania konfiguracją z wsparciem dla:
- Plików `.env` (przez godotenv) - różne pliki dla różnych środowisk
- Zmiennych środowiskowych systemowych

### Szybki start

1. Skopiuj przykładowy plik konfiguracyjny:
   ```bash
   cp .env.example .env
   ```

2. Dostosuj wartości w `.env` do swoich potrzeb:
   ```bash
   # Edytuj .env
   nano .env
   ```

3. Uruchom aplikację:
   ```bash
   go run cmd/api/main.go
   ```

### Środowiska

Ustaw zmienną `APP_ENV` aby przełączyć środowisko:
- `development` (domyślne) - lokalne środowisko deweloperskie
- `test` - środowisko testowe
- `production` - środowisko produkcyjne

Przykład:
```bash
APP_ENV=production go run cmd/api/main.go
```

### Precedencja wartości

Konfiguracja jest ładowana w kolejności (ostatnia wygrywa):

1. Wartości domyślne w kodzie
2. Plik `.env.{APP_ENV}` (np. `.env.development`) jeśli istnieje
3. Plik `.env` (fallback)
4. Zmienne środowiskowe systemowe

**Uwaga:** Zmienne środowiskowe systemowe mają najwyższy priorytet i nadpisują wartości z plików `.env`.

### Pliki konfiguracyjne

- `.env.example` - szablon konfiguracji (commitowany do repo)
- `.env` - domyślna konfiguracja (w .gitignore)
- `.env.development` - specyficzne dla `APP_ENV=development`
- `.env.test` - specyficzne dla `APP_ENV=test`
- `.env.production` - specyficzne dla `APP_ENV=production`

**Przykład użycia:**
```bash
# Załaduje .env.production
APP_ENV=production go run cmd/api/main.go

# Załaduje .env.development (lub .env jako fallback jeśli nie istnieje)
APP_ENV=development go run cmd/api/main.go
```

### Zmienne środowiskowe

| Zmienna | Opis | Domyślna wartość | Wymagana |
|---------|------|------------------|----------|
| `APP_ENV` | Środowisko aplikacji | `development` | Nie |
| `SERVER_ADDRESS` | Adres serwera HTTP | `:8080` | Nie |
| `DATABASE_URL` | URL połączenia do PostgreSQL | - | Tak |
| `JWT_SECRET` | Sekret do podpisywania JWT | - | Tak (production) |
| `LOG_LEVEL` | Poziom logowania | `info` | Nie |

### Walidacja konfiguracji

W środowisku `production` aplikacja wymaga:
- Ustawienia `JWT_SECRET` (nie może być wartości domyślnej)
- Ustawienia `DATABASE_URL`

W środowiskach `development` i `test` brak wymaganych wartości spowoduje jedynie ostrzeżenia.

## Uruchamianie

```bash
# Development (domyślnie)
go run cmd/api/main.go

# Production
APP_ENV=production go run cmd/api/main.go

# Z własnym plikiem .env
go run cmd/api/main.go
```

## Development Tools

Projekt korzysta z narzędzi automatycznego formatowania, lintingu i analizy statycznej kodu.

### Instalacja narzędzi

Użyj przygotowanego skryptu instalacyjnego:
```bash
# W katalogu backend
bash .golangci-install.sh
```

Lub zainstaluj ręcznie narzędzia używając Makefile:
```bash
make install-tools
```

### Dostępne komendy

Wszystkie komendy są dostępne przez `Makefile`. Zobacz listę dostępnych komend:

```bash
make help
```

#### Główne komendy:

```bash
# Przeanalizuj kod linterami
make lint

# Automatyczne naprawy
make lint-fix

# Formatuj kod
make fmt

# Sprawdź kod go vet
make vet

# Uruchom testy
make test

# Testy z raportem pokrycia
make test-coverage

# Skanowanie bezpieczeństwa
make security

# Wszystkie sprawdzenia (lint + vet + test)
make check

# Pełny workflow (fmt + lint + vet + test)
make all

# CI workflow (lint + vet + test)
make ci

# Build aplikacji
make build

# Uruchom aplikację
make run
```

### Konfiguracja VS Code

Workspace jest prekonfigurowany z następującymi ustawieniami:

- ✅ Automatyczne formatowanie przy zapisie
- ✅ Organizacja importów przy zapisie
- ✅ Linting z golangci-lint
- ✅ Vet check przy zapisie
- ✅ Linting tylko w edytowanych plikach

Aby korzystać z pełnej funkcjonalności, zainstaluj rozszerzenia:
- [Go extension for VS Code](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [golangci-lint for VS Code](https://marketplace.visualstudio.com/items?itemName=Golang.go)

### Lintery

Projekt używa **golangci-lint** - meta-lintera agregującego ponad 50 linterów. Konfiguracja znajduje się w `.golangci.yml`.

#### Aktywne lintery:

**Krytyczne:**
- `errcheck` - nieobsługiwane błędy
- `staticcheck` - zaawansowana analiza statyczna
- `govet` - problemy strukturalne
- `ineffassign` - nieużywane przypisania
- `unused` - nieużywany kod
- `typecheck` - sprawdzanie typów
- `gofmt` / `goimports` - formatowanie

**Jakość kodu:**
- `gocritic` - krytyczna analiza kodu
- `goconst` - wyszukiwanie magic strings
- `gofumpt` - zaawansowane formatowanie
- `misspell` - literówki
- `unparam` - nieużywane parametry
- `unconvert` - zbędne konwersje

**Złożoność i duplikacja:**
- `gocyclo` - złożoność cyklomatyczna
- `dupl` - duplikacja kodu
- `funlen` - długość funkcji

**Bezpieczeństwo:**
- `gosec` - skanowanie bezpieczeństwa

### CI/CD Integration

GitHub Actions automatycznie uruchamia:
- ✅ Linting z `golangci-lint`
- ✅ Static analysis z `go vet`
- ✅ Testy unitowe z raportem pokrycia
- ✅ Security scan z `gosec`
- ✅ Build verification

Workflowy są uruchamiane przy każdym PR i push do `main`/`develop`.

### Przed każdym commit

Zalecane jest uruchomienie pełnego check przed commit:
```bash
make check
```

Jeśli chcesz sformatować kod i naprawić automatycznie naprawialne problemy:
```bash
make all
```

### Troubleshooting

**Problem:** `golangci-lint: command not found`
```bash
# Zainstaluj golangci-lint
make install-tools

# Lub ręcznie
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

**Problem:** VS Code nie pokazuje błędów z golangci-lint
- Zainstaluj rozszerzenie "golangci-lint for VS Code"
- Sprawdź czy `golangci-lint` jest w `PATH`
- Zrestartuj VS Code

**Problem:** Makefile commands not found
```bash
# Upewnij się, że jesteś w katalogu backend
cd backend

# Sprawdź dostępność make
which make

# macOS: brew install make
# Linux: sudo apt-get install build-essential
```

