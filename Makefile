include .env

SHELL:=/bin/bash

.PHONY: migrate
migrate:
	migrate -path ./cmd/migrate/migrations -database ${DB_SRC} up

.PHONY: backend
backend:
	go run ./cmd/social

