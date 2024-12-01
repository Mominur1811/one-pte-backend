MAIN:=./cmd
TARGET:=onepte
SERVER_CMD:=./${TARGET}

# Just for documentation
MIGRATION_CREATE_EXAMPLE:=migrate create -ext sql -seq -dir ./migrations alter_default_order_acceptance_amount_limit_at_rider_status_table

run-server:
	${SERVER_CMD}

tidy:
	go mod tidy

install-dev-deps:
	go install github.com/air-verse/air@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

install-deps:
	go mod download

prepare: install-dev-deps install-deps tidy

dev:
	air

build: install-deps
	go build -o ${TARGET} ${MAIN}

start: build run-server

# Format using both gofumpt and golines
fmt:
	@echo "Formatting Go files with gofumpt and golines..."
	@gofumpt -w . && golines -w .
	@echo "All Go files formatted!"