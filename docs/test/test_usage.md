# Testing Models - Usage Guide

This document provides instructions on how to test the models in the Tranzure application, both on Windows and WSL Ubuntu 24.04.

## Prerequisites

- Go 1.22 or later installed
- For WSL testing: Ubuntu 24.04 LTS with WSL
- Git (for cloning the repository if needed)

## Testing on Windows

### Using Go Commands Directly

1. Open Command Prompt or PowerShell
2. Navigate to the project directory:
   ```
   cd C:\Users\KENTO\Programs\tranzure
   ```
3. Run the tests:
   ```
   go test -v ./internal/models/tests/...
   ```

### Using Make (requires Make installed on Windows)

1. Open Command Prompt or PowerShell
2. Navigate to the project directory:
   ```
   cd C:\Users\KENTO\Programs\tranzure
   ```
3. Run the model tests:
   ```
   make test-models
   ```

## Testing on WSL Ubuntu 24.04

### Using the Test Script

1. Open WSL terminal
2. Navigate to the mounted project directory:
   ```
   cd /mnt/c/Users/KENTO/Programs/tranzure
   ```
3. Make the test script executable (if not already):
   ```
   chmod +x scripts/test_models.sh
   ```
4. Run the test script:
   ```
   ./scripts/test_models.sh
   ```

### Using Make

1. Open WSL terminal
2. Navigate to the mounted project directory:
   ```
   cd /mnt/c/Users/KENTO/Programs/tranzure
   ```
3. Run the model tests:
   ```
   make test-models
   ```

## What to Expect

When tests run successfully, you should see output similar to:

```
===== Tranzure Model Tests =====
Using Go version: go version go1.22.2 linux/amd64
Project root: /mnt/c/Users/KENTO/Programs/tranzure
Downloading dependencies...
Running model tests...
=== RUN   TestWalletValidation
--- PASS: TestWalletValidation (0.00s)
=== RUN   TestUserValidation
--- PASS: TestUserValidation (0.00s)
// ... more tests ...
PASS
ok      github.com/kento/tranzure/internal/models/tests   0.032s
Tests completed.
```

## Troubleshooting

### Common Issues

1. **Missing Dependencies**
   If you encounter dependency errors, try:
   ```
   go mod tidy
   ```

   If you see an error like:
   ```
   missing go.sum entry for module providing package gorm.io/datatypes
   ```
   Run the specific command mentioned in the error message:
   ```
   go get -t github.com/kento/tranzure/internal/models/tests
   ```

   For stubborn dependency issues, try:
   ```
   # Clean the module cache
   go clean -modcache

   # Reinitialize dependencies
   go mod tidy
   go mod download
   go mod verify
   ```

2. **Permission Denied in WSL**
   Ensure the test script is executable:
   ```
   chmod +x scripts/test_models.sh
   ```

3. **Test Script Not Found**
   Make sure you're in the project root directory when running the script.

4. **Failed Tests**
   - Read the error messages carefully for clues about what's failing
   - Check model validation rules
   - Verify that your test data matches expected formats and constraints

### Getting Help

If you continue to have issues:
- Check the GitHub repo issues section
- Open a new issue with detailed information about the error
- Contact the project maintainer
