# Define main and target
MAIN := ./cmd
TARGET := onepte
SERVER_CMD := ./${TARGET}

# Windows-specific variables
TARGET_WIN := $(TARGET).exe
SERVER_CMD_WIN := .\$(TARGET).exe
MAKE_DEV_WIN := dev_win

# Linux/macOS specific variables
MAKE_DEV_LINUX := dev

# Migration example (for documentation)
MIGRATION_CREATE_EXAMPLE := migrate create -ext sql -seq -dir ./migrations alter_default_order_acceptance_amount_limit_at_rider_status_table

# Run server for Linux/macOS
run-server-linux:
	${SERVER_CMD}

# Run server for Windows
run-server-win:
	${SERVER_CMD_WIN}

# Tidy Go modules
tidy:
	go mod tidy

# Install development dependencies
install-dev-deps:
	go install github.com/cosmtrek/air@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install general dependencies
install-deps:
	go mod download

# Prepare project (install dev deps, install deps, tidy)
prepare: install-dev-deps install-deps tidy

# Development server for Linux/macOS
dev:
	air -c .air.toml

# Development server for Windows
dev_win:
	air -c .air.win.toml

# Build project for Linux/macOS
build-linux: install-deps
	go build -o ${TARGET} ${MAIN}

# Build project for Windows
build-win: install-deps
	go build -o ${TARGET_WIN} ${MAIN}

# Start the project (build + run server for Linux/macOS)
start-linux: build-linux run-server-linux

# Start the project (build + run server for Windows)
start-win: build-win run-server-win

# Format Go files using gofumpt and golines
fmt:
	@echo "Formatting Go files with gofumpt and golines..."
	@gofumpt -w . && golines -w .
	@echo "All Go files formatted!"
