# Testing Models on WSL Ubuntu 24.04

This guide explains how to run the model tests in WSL Ubuntu 24.04 LTS.

## Prerequisites

- Windows Subsystem for Linux (WSL) with Ubuntu 24.04 LTS installed
- Basic familiarity with terminal commands

## Setting Up

1. Open WSL terminal (Ubuntu 24.04)

2. Navigate to the project directory:
   ```bash
   cd /mnt/c/Users/KENTO/Programs/tranzure
   ```

3. Make the test script executable:
   ```bash
   chmod +x scripts/test_models.sh
   ```

## Running Tests

### Option 1: Using the script directly

Run the script:
```bash
./scripts/test_models.sh
```

### Option 2: Using Make (easier)

If you prefer using Make:
```bash
make test-models
```

## What the Script Does

The script will:
- Check if Go is installed and install it if needed
- Set up the Go environment
- Download dependencies
- Run all model tests in verbose mode

## Troubleshooting

If you encounter errors:

1. **Permission denied**: Make sure the script is executable
   ```bash
   chmod +x scripts/test_models.sh
   ```

2. **Go module issues**: Try cleaning the Go module cache
   ```bash
   go clean -modcache
   ```

3. **WSL path issues**: Ensure the project is accessible from WSL
   ```bash
   ls -la /mnt/c/Users/KENTO/Programs/tranzure
   ```

4. **Dependencies**: If there are issues with dependencies, try:
   ```bash
   go mod tidy
   go mod verify
   ```
