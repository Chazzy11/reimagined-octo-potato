.PHONY: build run test clean fmt install help

# Default target
all: build

# Build the binary
build:
	@echo "Building dark labyrinth..."
	@go build -o darklabyrinth .

# Run the game
run: build
	@./darklabyrinth

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f darklabyrinth
	@go clean

# Install to GOPATH/bin
install:
	@echo "Installing..."
	@go install

# Display help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  run      - Build and run the game"
	@echo "  test     - Run tests"
	@echo "  fmt      - Format code with gofmt"
	@echo "  clean    - Remove build artifacts"
	@echo "  install  - Install to GOPATH/bin"
	@echo "  help     - Show this help message"