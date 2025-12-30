APP_NAME=postgress-debezium-kafka-app

.PHONY: build run test tidy

build:
	go build -o bin/${APP_NAME} ./cmd/server

run: build
	./bin/${APP_NAME}

test:
	go test ./...

tidy:
	go mod tidy

# Lint
lint:
	# Install with `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.59.0`
	golangci-lint run ./...


# Docker helpers
docker-build:
	docker build -t ${APP_NAME}:local .

compose-up:
	docker compose up -d --build

compose-down:
	docker compose down -v

# Run DB migrations using the local Go binary (reads DATABASE_URL from env or .env)
migrate:
	go run ./cmd/migrate

# Run migrations inside the app container (uses compose services)
compose-migrate:
	docker compose run --rm app /app/migrate
