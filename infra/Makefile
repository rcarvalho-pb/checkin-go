APP_NAME=checkin
CMD_DIR=cmd/server
MIGRATIONS_DIR=migrations
DB_TYPE=postgres
DSN=postgres://postgres:postgres@localhost:5432/checkin?sslmode=disable
SECRET=very-secret

.PHONY: dev build run migrate up down ps logs

dev:
	@go tool templ generate
	@go run ../${CMD_DIR}

build:
	@go build -o ../bin/${APP_NAME} ../${CMD_DIR}

run: stop build
	@go tool templ generate
	@export DSN=${DB_DSN} && ../bin/${APP_NAME} &

stop:
	@-pkill -SIGTERM -f "${APP_NAME}"

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
