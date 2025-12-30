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

## Docker Compose

You can run the app and a Postgres instance locally with Docker Compose:

1. Copy `.env.example` to `.env` and edit if needed (DATABASE_URL points to the `db` service)
2. Bring up services: `docker compose up -d --build`

The app runs migrations on start when `MIGRATE_ON_START=true` in `.env`.

### Database web UI

An Adminer service is included for quick DB access. After bringing up the stack you can open Adminer at:

- http://localhost:${ADMINER_PORT:-8081}

Use the credentials from `.env` (defaults are `postgres` / `postgres`) and connect to host `db` (the compose service name).

### Debezium (change data capture)

This repository includes a simple Debezium stack (Zookeeper, Kafka, Kafka Connect) to capture Postgres changes and stream them to Kafka.

What's included

- `kafka` (KRaft mode) and `connect` services in `docker-compose.yml` (no Zookeeper)
- DB init script that creates a `debezium` replication user and publication `dbz_publication`
- An automatic connector registration script that registers a Postgres connector with Kafka Connect

How to run

1. Copy `.env.example` to `.env` and ensure `MIGRATE_ON_START=true` and DB vars are set.
2. Start the stack: `docker compose up -d --build`
3. After services are up, register connector (the `connect-init` service tries to register the connector automatically; you may re-run it using `docker compose run --rm connect-init`).
4. Kafka Connect UI: http://localhost:8083
5. Adminer: http://localhost:${ADMINER_PORT:-8081}

Notes

- Postgres is started with `wal_level=logical` and additional replication settings so Debezium can stream changes.
- Connector config is in `docker/debezium/postgres-connector.json` and can be customized.

