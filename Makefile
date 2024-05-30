ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable
endif


PIS_MIGRATION_FOLDER=$(CURDIR)/products-info-service/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(PIS_MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(PIS_MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(PIS_MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: full-compose-up
full-compose-up:
	docker compose build
	docker compose up -d postgres
	docker compose up -d redis
	docker compose up -d circuit-breaker
	docker compose up -d products-info-service
	docker compose up -d create-order-service

.PHONY: without-cb-compose-up
without-cb-compose-up:
	docker compose build
	docker compose up -d postgres
	docker compose up -d redis
	docker compose up -d products-info-service
	docker compose up -d create-order-service

.PHONY: without-cos-compose-up
without-cos-compose-up:
	docker compose build
	docker compose up -d postgres
	docker compose up -d redis
	docker compose up -d circuit-breaker
	docker compose up -d products-info-service

.PHONY: compose-rm
compose-rm:
	docker compose down

.PHONY up-all:
up-all:
	make full-compose-up
	make migration-up
	go run add-products-service/cmd/main.go

.PHONY up-storages:
up-storages:
	docker compose build
	docker compose up -d postgres
	docker compose up -d redis
	make migration-up

