APP_NAME := my-fiber-app
BUILD_DIR := ./build
GO_FILES := $(shell find . -type f -name '*.go')

# Default target
.PHONY: all
all: build

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Generate swagger docs
swag:
	@echo "Generating swagger docs..."
	@swag init

# Wire
wire:
	@echo "Running wire..."
	@wire

# Run the application
.PHONY: run
run: wire swag
	@echo "Running $(APP_NAME)..."
	@go run main.go

# Execute tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@gofmt -w $(GO_FILES)

# Lint the code
.PHONY: lint
lint:
	@echo "Linting code..."
	@golangci-lint run

# Build the application
.PHONY: build
build: deps wire swag
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME)

# Build the appliation docker image
.PHONY: docker-build
docker-build:
	@echo "Building docker image..."
	@docker build -t $(APP_NAME):latest .

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Help
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make deps      Install dependencies"
	@echo "  make run       Run the application"
	@echo "  make test      Run tests"
	@echo "  make fmt       Format code"
	@echo "  make lint      Lint the code"
	@echo "  make build     Build the application"
	@echo "  make clean     Clean build artifacts"