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

# Docker helpers
docker-build:
	docker build -t ${APP_NAME}:local .

compose-up:
	docker compose up -d --build

compose-down:
	docker compose down -v
