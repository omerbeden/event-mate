SHELL := /bin/bash

# Network name
NETWORK_NAME := event-mate-network

# Database configurations
DB_PASSWORD := password
DB_NAME := test
DB_PORT := 5432
DB_USER := postgres
DB_CONTAINER_NAME := dev-postgres

# Migration configurations
MIGRATIONS_PATH := /Users/omerbeden/golang_workspace/event-mate/backend/tatooine/infra/databasemigrations
DATABASE_URL := "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"

	
.PHONY: help network-up db-up db-down migrate-up integration-test-setup


# Display help for each command
help:
	@echo "Usage: make [COMMAND]"
	@echo "Commands:"
	@echo "  network-up           Create Docker network if it doesn't exist"
	@echo "  db-up                Start the PostgreSQL container"
	@echo "  db-down              Stop and remove the PostgreSQL container"
	@echo "  migrate-up           Run database migrations"
	@echo "  integration-test-setup Prepare environment for integration testing (starts db and runs migrations)"

# Create Docker network if it doesn't exist
network-up:
	@if [ -z $$(docker network ls --filter name=^${NETWORK_NAME}$$ --format="{{ .Name }}") ]; then \
		echo "Creating Docker network '${NETWORK_NAME}'..."; \
		docker network create ${NETWORK_NAME}; \
	else \
		echo "Docker network '${NETWORK_NAME}' already exists."; \
	fi

# Start the PostgreSQL container
db-up: network-up
	docker run --rm -d \
		--name $(DB_CONTAINER_NAME) \
		-p $(DB_PORT):5432 \
		--network $(NETWORK_NAME) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		postgres

# Stop the PostgreSQL container
db-down:
	@docker stop $(DB_CONTAINER_NAME) || true

# Run database migrations
migrate-up:	
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URL) -verbose up

# Target to run before integration tests
integration-test-setup: db-up migrate-up
