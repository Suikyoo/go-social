include .env

SHELL:=/bin/bash
BACKEND_PATH:=./cmd/social
FRONTEND_PATH:=./web/social

.PHONY: migrate-up
migrate-up:
	migrate -path ./cmd/migrate/migrations -database ${DB_SRC} up

.PHONY: migrate-down
migrate-down:
	migrate -path ./cmd/migrate/migrations -database ${DB_SRC} down


.PHONY: backend
backend:
	go run $(BACKEND_PATH)

.PHONY: frontend
frontend:
	cd $(FRONTEND_PATH) && npm install && npm run dev

