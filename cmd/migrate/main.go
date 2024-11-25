package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ushiradineth/cron-be/database"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: [up|down]\n")
		os.Exit(1)
	}

	command := os.Args[1]
	if err := run(command); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(command string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("Failed to load env")
	}

	db := database.New()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Failed to create Postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration",
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("Failed to create migration instance: %w", err)
	}

	switch command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		return fmt.Errorf("Unknown command: %s", command)
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("Migration command '%s' failed: %w", command, err)
	}

	fmt.Printf("Migration command '%s' executed successfully\n", command)
	return nil
}
