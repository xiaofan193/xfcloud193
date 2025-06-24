# =============================================================================
# PROJECT CONFIGURATION
# =============================================================================
# Project metadata and core settings
PROJECT_NAME := myapp
VERSION := $(shell git describe --tags --always --dirty="-dev")
BUILD_DIR := bin
GO_PACKAGE_PATH := github.com/yourorg/$(PROJECT_NAME)

# API Server specific configuration
API_SERVER_DIR := cmd/apiserver
API_SERVER_BINARY := $(BUILD_DIR)/apiserver
API_SOURCE_FILES := $(shell find $(API_SERVER_DIR) -name '*.go')

# Go toolchain configuration
GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test -v -cover
GO_CLEAN := $(GO) clean
GO_MOD := $(GO) mod
GO_LINT := golangci-lint run
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Build flags for different environments
BUILD_FLAGS := -ldflags "-X $(GO_PACKAGE_PATH)/internal/version.Version=$(VERSION)"
ifeq ($(ENV), prod)
  BUILD_FLAGS += -ldflags "-w -s" # Strip debug symbols in production
  OUTPUT_DIR := $(BUILD_DIR)/production
else
  OUTPUT_DIR := $(BUILD_DIR)/development
endif

# =============================================================================
# BUILD TARGETS
# =============================================================================
# Default target: show help information
.DEFAULT_GOAL := help

# Build API server application
.PHONY: build-cmd/apiserver
build-cmd/apiserver: $(OUTPUT_DIR)/apiserver  ## Build API server application

$(OUTPUT_DIR)/apiserver: $(API_SOURCE_FILES)
	@echo "Building API server v$(VERSION) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(@D)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) $(BUILD_FLAGS) -o $@ ./$(API_SERVER_DIR)
	@echo "Build complete: $@"
	@du -h $@

# Alias for path-style invocation
.PHONY: cmd/apiserver
cmd/apiserver: build-cmd/apiserver

# Build all components
.PHONY: build-all
build-all: build-cmd/apiserver  ## Build all project components

# =============================================================================
# DEVELOPMENT WORKFLOW
# =============================================================================
# Run API server locally
.PHONY: run
run: build-cmd/apiserver  ## Run API server locally
	@echo "Starting API server..."
	@$(OUTPUT_DIR)/apiserver

# Run tests with coverage analysis
.PHONY: test
test:  ## Run tests with coverage
	@echo "Running tests..."
	$(GO_TEST) -coverprofile=cover.out ./...

# Generate HTML coverage report
.PHONY: coverage
coverage: test  ## Generate HTML coverage report
	@echo "Generating coverage report..."
	$(GO) tool cover -html=cover.out -o coverage.html
	@echo "Coverage report: coverage.html"

# =============================================================================
# CODE QUALITY
# =============================================================================
# Format Go code
.PHONY: fmt
fmt:  ## Format code with gofmt
	@echo "Formatting code..."
	$(GO) fmt ./...

# Run static code analysis
.PHONY: lint
lint:  ## Run static code analysis
	@echo "Running linters..."
	$(GO_LINT)

# Tidy Go module dependencies
.PHONY: tidy
tidy:  ## Tidy dependencies
	@echo "Tidying dependencies..."
	$(GO_MOD) tidy

# =============================================================================
# DEPLOYMENT & DISTRIBUTION
# =============================================================================
# Cross-compile for multiple platforms
.PHONY: cross-build
cross-build:  ## Build multi-platform binaries
	@echo "Cross-compiling API server..."
	@mkdir -p $(BUILD_DIR)/dist
	GOOS=linux GOARCH=amd64 $(MAKE) build-cmd/apiserver OUTPUT_DIR=$(BUILD_DIR)/dist
	GOOS=darwin GOARCH=arm64 $(MAKE) build-cmd/apiserver OUTPUT_DIR=$(BUILD_DIR)/dist
	GOOS=windows GOARCH=amd64 $(MAKE) build-cmd/apiserver OUTPUT_DIR=$(BUILD_DIR)/dist
	@echo "Binaries available in $(BUILD_DIR)/dist"

# Build Docker image
.PHONY: docker-build
docker-build:  ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(PROJECT_NAME)-apiserver:$(VERSION) -f $(API_SERVER_DIR)/Dockerfile .
	@docker image inspect $(PROJECT_NAME)-apiserver:$(VERSION) --format 'Image size: {{.Size}} bytes'

# =============================================================================
# MAINTENANCE
# =============================================================================
# Clean build artifacts
.PHONY: clean
clean:  ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR) cover.out coverage.html

# Generate dependency graph
.PHONY: deps
deps:  ## Generate dependency graph
	@echo "Generating dependency graph..."
	$(GO) mod graph | dot -Tpng -o dependency-graph.png

# =============================================================================
# HELP SYSTEM
# =============================================================================
# Display help information
.PHONY: help
help:  ## Show this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# =============================================================================
# UTILITY TARGETS
# =============================================================================
# Security scanning
.PHONY: security-scan
security-scan: build-cmd/apiserver  ## Run security scans
	@echo "Running security scans..."
	gosec ./...
	trivy fs --security-checks vuln $(OUTPUT_DIR)

# Generate protocol buffers
.PHONY: proto
proto:  ## Generate protocol buffer code
	@echo "Generating protocol buffer code..."
	protoc --go_out=. --go-grpc_out=. proto/*.proto

# Generate mock implementations
.PHONY: mocks
mocks:  ## Generate mock implementations
	@echo "Generating mocks..."
	mockery --all --case underscore --output ./mocks