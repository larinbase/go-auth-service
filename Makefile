.PHONY: build run test migrate up down

build:
	go build -o bin/auth-service ./cmd/auth

run:
	go run ./cmd/auth/main.go

test:
	go test -v ./...

up:
	docker-compose up --build -d

down:
	docker-compose down

migrate:
	@echo "Running database migrations..."
	lint:
	go vet ./...
	golint ./...

.DEFAULT_GOAL := build