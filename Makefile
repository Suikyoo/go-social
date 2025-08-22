include .env

SHELL:=/bin/bash

.PHONY: migrate
migrate:
	migrate -path ./cmd/migrate/migrations -database ${DB_SRC} up

.PHONY: create-user
create-user:
	curl --json '{"name": "suikyo", "password": "secret"}' http://localhost:8080/auth/user

