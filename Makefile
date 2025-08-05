# Variables
MAIN_PATH := ./cmd/app

# Run the application
.PHONY: run
run:
	go run $(MAIN_PATH)

# Run tests
.PHONY: test
test:
	go test -v ./...

# Install dependencies
.PHONY: deps
deps:
	go mod download

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy

# Show help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  run    - Run the application"
	@echo "  test   - Run tests"
	@echo "  deps   - Install dependencies"
	@echo "  tidy   - Tidy dependencies"
