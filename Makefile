include .env

SHELL:=/bin/bash
BACKEND_PATH:=./cmd/social
FRONTEND_PATH:=./web/social

.PHONY: migrate
migrate:
	migrate -path ./cmd/migrate/migrations -database ${DB_SRC} up

.PHONY: backend
backend:
	go run $(BACKEND_PATH)

.PHONY: frontend
frontend:
	cd $(FRONTEND_PATH) && npm install && npm run dev

