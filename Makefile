.PHONY: help build run test clean docker-up docker-down sqlc-generate migrate-up migrate-down sqlc-admin sqlc-merchant

# Default target
help:
	@echo "Available targets:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run tests"
	@echo "  clean          - Clean build artifacts"
	@echo "  docker-up      - Start Docker containers"
	@echo "  docker-down    - Stop Docker containers"
	@echo "  sqlc-generate  - Generate SQLC code"
	@echo "  migrate-up     - Run database migrations up"
	@echo "  migrate-down   - Run database migrations down"
	@echo "  dev            - Start development environment"

# Build the application
build:
	go build -o bin/main ./cmd/main.go

# Run the application
run:
	go run ./cmd/api/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Generate SQLC code (root default)
sqlc-generate:
	sqlc generate --file db/sqlc.yaml || true

# Generate Admin SQLC
sqlc-admin:
	sqlc generate -f db/sqlc/admin/sqlc.yaml

# Generate Merchant SQLC
sqlc-merchant:
	sqlc generate -f db/sqlc/merchant/sqlc.yaml

# Run database migrations up
migrate-up:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/users_db?sslmode=disable" up

# Run database migrations down
migrate-down:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/users_db?sslmode=disable" down

# Development environment
dev: docker-up
	@echo "Waiting for database to be ready..."
	@sleep 5
	@echo "Generating SQLC code..."
	@make sqlc-generate
	@echo "Starting application..."
	@make run
