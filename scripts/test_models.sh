#!/bin/bash

# Script to run model tests on WSL Ubuntu
# Usage: ./scripts/test_models.sh [specific_test]
#
# Examples:
#   ./scripts/test_models.sh                    # Run all model tests
#   ./scripts/test_models.sh TestWalletValidation  # Run a specific test

# Exit on error
set -e

echo "===== Tranzure Model Tests ====="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."
    sudo apt-get update
    sudo apt-get install -y wget
    wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
    rm go1.22.2.linux-amd64.tar.gz

    # Set up Go environment variables
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
fi

# Print Go version
echo "Using Go version: $(go version)"

# Set environment variables for testing
export GO111MODULE=on

# Get directory of this script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "Project root: $PROJECT_ROOT"
cd "$PROJECT_ROOT"

# Ensure dependencies are downloaded and fixed
echo "Fixing dependencies..."
go mod tidy
go get -t ./internal/models/tests/...
go mod download

# Run model tests
echo "Running model tests..."
cd "$PROJECT_ROOT/internal/models/tests"

# Check if a specific test was specified
if [ -n "$1" ]; then
    echo "Running specific test: $1"
    go test -v -run "$1"
else
    # Run all tests with verbose output
    go test -v ./...
fi

# Check test status
if [ $? -eq 0 ]; then
    echo "✅ Tests completed successfully."
else
    echo "❌ Tests failed."
    exit 1
fi
