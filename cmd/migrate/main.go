package main

import (
	"github.com/dfryer1193/werewolf/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/dfryer1193/werewolf/internal/config"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.LoadConfig()

	pgdb, err := db.NewPostgresDB(cfg.DBConnString)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	conn, err := postgres.WithInstance(pgdb.Pool, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create database instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", db)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("Please provide a migration command")
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Could not apply migrations: %v", err)
		}
		log.Println("Migrations applied")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Could not rollback migrations: %v", err)
		}
		log.Println("Migrations rolled back")
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
