package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // 👈 use postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Build DSN for PostgreSQL
	// Important: sslmode=disable if you run local dev without SSL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname)

	// Migrations path
	migrationsPath := "file://src/migrations"

	// Init migrate
	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Default: up
	direction := "up"
	if len(os.Args) > 1 {
		direction = os.Args[1]
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Migration up complete")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Migration down complete")
	default:
		fmt.Println("Usage: go run ./cmd/migrate.go [up|down]")
	}
}

