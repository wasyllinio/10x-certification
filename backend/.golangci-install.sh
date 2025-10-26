#!/bin/bash

# Installation script for golangci-lint and development tools
# Usage: bash .golangci-install.sh

set -e

GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${CYAN}==================================================${NC}"
echo -e "${CYAN}Installing Go Development Tools${NC}"
echo -e "${CYAN}==================================================${NC}"
echo

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed. Please install Go first.${NC}"
    exit 1
fi

echo -e "${GREEN}Go version: $(go version)${NC}"
echo

# Function to install a tool if not already installed
install_tool() {
    local tool_name=$1
    local install_command=$2
    
    if command -v "$tool_name" &> /dev/null; then
        echo -e "${YELLOW}✓ $tool_name is already installed${NC}"
        return
    fi
    
    echo -e "${CYAN}Installing $tool_name...${NC}"
    eval "$install_command"
    
    if command -v "$tool_name" &> /dev/null; then
        echo -e "${GREEN}✓ $tool_name installed successfully${NC}"
    else
        echo -e "${RED}✗ Failed to install $tool_name${NC}"
        return 1
    fi
}

# Install tools
echo -e "${CYAN}Installing development tools...${NC}"
echo

install_tool "golangci-lint" "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/main/install.sh | sh -s -- -b $(go env GOPATH)/bin"
install_tool "gosec" "go install github.com/securego/gosec/v2/cmd/gosec@latest"
install_tool "goimports" "go install golang.org/x/tools/cmd/goimports@latest"

echo
echo -e "${CYAN}==================================================${NC}"
echo -e "${GREEN}Installation complete!${NC}"
echo -e "${CYAN}==================================================${NC}"
echo
echo -e "${CYAN}Installed tools:${NC}"
echo -e "${GREEN}  • golangci-lint${NC}"
echo -e "${GREEN}  • gosec (security scanner)${NC}"
echo -e "${GREEN}  • goimports (import organizer)${NC}"
echo
echo -e "${CYAN}Quick start:${NC}"
echo -e "  Run \`make lint\` to check your code"
echo -e "  Run \`make lint-fix\` to auto-fix issues"
echo -e "  Run \`make fmt\` to format your code"
echo -e "  Run \`make check\` to run all checks"
echo
echo -e "${CYAN}For more commands, run: make help${NC}"

