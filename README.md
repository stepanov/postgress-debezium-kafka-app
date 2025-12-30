# postgress-debezium-kafka-app

Scaffolded Go API with PostgreSQL CRUD using models and repositories.

## Quickstart

1. Set env var: `DATABASE_URL` (e.g. `postgresql://user:pass@localhost:5432/dbname`)
2. Run migration: `go run ./cmd/migrate`
3. Start server: `go run ./cmd/server` (or `make run`) — default port `8080`.

## Endpoints

- POST /users          (create user)
- GET  /users/{id}     (get single user)
- GET  /users          (list users)
- PUT  /users/{id}     (update user)
- DELETE /users/{id}   (delete user)

## Project layout

- `internal/model` – domain models
- `internal/repository` – repository interfaces
- `internal/repository/postgres` – Postgres implementation
- `internal/handlers` – HTTP handlers
- `pkg/db` – DB connection factory
- `cmd/server` – HTTP server
- `cmd/migrate` – simple SQL migration

Enjoy!"}