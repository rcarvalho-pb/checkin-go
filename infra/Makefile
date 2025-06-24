APP_NAME=checkin
CMD_DIR=cmd/server
MIGRATIONS_DIR=migrations
DB_DSN=postgres://postgres:password@localhost:5432/checkin?sslmode=disable

.PHONY: dev build run migrate up down ps logs

dev:
	templ generate
	go run ../${CMD_DIR}

build:
	go build -o ../bin/${APP_NAME} ../${CMD_DIR}

run: build
	export DSN=${DB_DSN} && ../bin/${APP_NAME}

migrate:
	go run ../${CMD_DIR}

up:
	docker compose up -d

down:
	docker compose down

ps:
	docker compose ps

logs:
	docker compose logs -f
