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
