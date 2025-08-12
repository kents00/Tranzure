# Tranzure Makefile

.PHONY: test test-models clean build run docker-up docker-down

# Default target
all: build

# Build the application
build:
	go build -v ./...

# Run the application
run: build
	./main

# Clean build artifacts
clean:
	go clean
	rm -f main

# Run all tests
test:
	go test -v ./...

# Run model tests specifically (works in both Windows and WSL)
test-models:
	@if command -v bash > /dev/null 2>&1; then \
		bash scripts/test_models.sh; \
	else \
		go test -v ./internal/models/tests/...; \
	fi

# Test database functionality
test-database:
	@if command -v bash > /dev/null 2>&1; then \
		bash scripts/test_database.sh; \
	else \
		echo "Database testing requires bash and Docker"; \
	fi

# Run all tests including database
test-all: test-models test-database

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  run           - Build and run the application"
	@echo "  test          - Run all tests"
	@echo "  test-models   - Run model tests"
	@echo "  test-database - Run database tests"
	@echo "  test-all      - Run all tests including database"
	@echo "  clean         - Clean build artifacts"
	@echo "  docker-up     - Start Docker containers"
	@echo "  docker-down   - Stop Docker containers"
