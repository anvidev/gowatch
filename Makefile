.PHONY: all build build-server build-watcher run-server run-watcher clean build-watcher-linux build-watcher-windows build-watcher-mac

# Define the output binaries
SERVER_BINARY := bin/server
WATCHER_BINARY := bin/watcher

# Default target
all: build

# Build both server and watcher binaries
build: build-server build-watcher

# Build the server binary
build-server:
	@echo "Building server..."
	@go build -o $(SERVER_BINARY) ./cmd/server

# Build the watcher binary
build-watcher:
	@echo "Building watcher..."
	@go build -o $(WATCHER_BINARY) ./cmd/watcher

# Build the watcher binary for Linux
build-watcher-linux:
	@echo "Building watcher for Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(WATCHER_BINARY)-linux ./cmd/watcher

# Build the watcher binary for Windows
build-watcher-windows:
	@echo "Building watcher for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(WATCHER_BINARY)-windows.exe ./cmd/watcher

# Build the watcher binary for macOS
build-watcher-mac:
	@echo "Building watcher for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(WATCHER_BINARY)-mac ./cmd/watcher

# Run the server binary
run-server: build-server
	@echo "Running server..."
	@./$(SERVER_BINARY)

# Run the watcher binary
run-watcher: build-watcher
	@echo "Running watcher..."
	@./$(WATCHER_BINARY)