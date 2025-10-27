.PHONY: help lint lint-fix fmt vet test test-coverage check security install-tools clean build run

# Colors for output
CYAN := \033[36m
GREEN := \033[32m
YELLOW := \033[33m
RED := \033[31m
NC := \033[0m # No Color

# Default target
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo "$(CYAN)Available commands:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2}'

lint: ## Run golangci-lint analysis
	@echo "$(CYAN)Running golangci-lint...$(NC)"
	@golangci-lint run

lint-fix: ## Run golangci-lint with auto-fix
	@echo "$(CYAN)Running golangci-lint with auto-fix...$(NC)"
	@golangci-lint run --fix

fmt: ## Format code with gofmt and goimports
	@echo "$(CYAN)Formatting code...$(NC)"
	@go fmt ./...
	@gofmt -w .
	@test -z $$(gofmt -l .) || (echo "$(RED)Code is not properly formatted. Run 'make fmt'$(NC)" && exit 1)

vet: ## Run go vet
	@echo "$(CYAN)Running go vet...$(NC)"
	@go vet ./...

test: ## Run tests
	@echo "$(CYAN)Running tests...$(NC)"
	@go test -v -race -coverprofile=coverage.out ./...

test-coverage: ## Run tests with coverage report
	@echo "$(CYAN)Running tests with coverage...$(NC)"
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

security: ## Run security scan with gosec
	@echo "$(CYAN)Running security scan...$(NC)"
	@which gosec > /dev/null || (echo "$(YELLOW)gosec not found. Installing...$(NC)" && go install github.com/securego/gosec/v2/cmd/gosec@latest)
	@gosec -severity medium -confidence medium ./...
	@golangci-lint run --enable gosec

check: lint vet test ## Run all checks (lint, vet, test)
	@echo "$(GREEN)All checks passed!$(NC)"

install-tools: ## Install development tools
	@echo "$(CYAN)Installing development tools...$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "$(GREEN)Tools installed successfully!$(NC)"

clean: ## Clean build artifacts
	@echo "$(CYAN)Cleaning...$(NC)"
	@rm -f coverage.out coverage.html
	@go clean ./...

build: ## Build the application
	@echo "$(CYAN)Building application...$(NC)"
	@go build -o bin/api ./cmd/api/main.go

run: ## Run the application
	@echo "$(CYAN)Running application...$(NC)"
	@go run -tags viper_bind_struct ./cmd/api/main.go

# Aliases for common workflows
ci: check ## Run CI checks (lint + vet + test)
all: fmt lint vet test ## Run everything (fmt + lint + vet + test)

