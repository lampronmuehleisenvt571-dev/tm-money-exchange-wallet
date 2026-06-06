.PHONY: help build test run dev-up dev-down migrate lint format clean

APP_NAME=tm-money-exchange
DOCKER_COMPOSE_FILE=deployment/docker-compose.yml

## help: Display this help message
help:
	@echo "TM Money Exchange Wallet - Makefile Commands"
	@echo ""
	@echo "Development:"
	@echo "  make dev-up              Start development environment (Docker Compose)"
	@echo "  make dev-down            Stop development environment"
	@echo "  make dev-run             Run application locally"
	@echo ""
	@echo "Testing:"
	@echo "  make test                Run all tests"
	@echo "  make test-coverage       Run tests with coverage"
	@echo "  make lint                Run linters"
	@echo ""
	@echo "Database:"
	@echo "  make db-migrate          Run database migrations"
	@echo "  make db-seed             Seed database with test data"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean               Clean build artifacts"

## dev-up: Start development environment
dev-up:
	@echo "Starting development environment..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

## dev-down: Stop development environment
dev-down:
	@echo "Stopping development environment..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

## dev-run: Run application locally
dev-run:
	@echo "Running application..."
	go run ./src/main.go

## test: Run all tests
test:
	@echo "Running tests..."
	go test -v -race ./...

## test-coverage: Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

## lint: Run linters
lint:
	@echo "Running linters..."
	golangci-lint run ./...

## format: Format code
format:
	@echo "Formatting code..."
	go fmt ./...

## clean: Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

.DEFAULT_GOAL := help
