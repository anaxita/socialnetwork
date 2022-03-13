.PHONY: all format lint build run migrate-install migrate-new migrate-up migrate-down golangci-lint up-local down-local gen

.DEFAULT_GOAL: format lint

-include deployment/.env

MIGRATE_TOOL_URL := https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz
LINTER_URL := https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh

all: format lint build

format:
	gofmt -s -w . && \
	go vet ./... && \
	go mod tidy

build:
	go build -o ./build/synergycommunity ./cmd/community

run:
	go run ./cmd/community

test:
	go clean -testcache && \
	go test -v ./...
# Docker
up-local:
	cd deployment && \
	docker-compose -f docker-compose.yml --env-file .env up -d

down-local:
	cd deployment && \
    docker-compose -f docker-compose.yml down

gen:
	go generate ./...

# Linter
lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0

lint: lint-install
	golangci-lint run

lint-fix: lint-install
	golangci-lint run --fix --path-prefix ''

# Migrations
migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1

migrate-new: migrate-install
	migrate create -ext sql -dir ./migrations "$(name)"


migrate-up: migrate-install
	migrate -database="${DB_SCHEME}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations up

migrate-down: migrate-install
	migrate -database="${DB_SCHEME}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations down 1

migrate-drop: migrate-install
	migrate -database="${DB_SCHEME}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations down


migrate-version: migrate-install
	migrate -database="${DB_SCHEME}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations force ${version}

seed:
	go run ./cmd/faker -users 20 -tags 20 -groups 20 -posts 20 -subs 60

migrate-rebase: migrate-drop migrate-up seed
