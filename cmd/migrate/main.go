package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/stepanov/postgress-debezium-kafka-app/pkg/db"
)

func main() {
	// load .env if present
	_ = godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pool, err := db.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}
	defer pool.Close()

	_, err = pool.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		id uuid PRIMARY KEY,
		name text NOT NULL,
		email text NOT NULL,
		created_at timestamptz NOT NULL DEFAULT now()
	);
	`)
	if err != nil {
		log.Fatalf("run migration: %v", err)
	}
	log.Println("migration applied")
}
