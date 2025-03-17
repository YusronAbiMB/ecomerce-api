package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/YusronAbi/ecomerce-api/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg config.Config) error {
	pass := url.QueryEscape(cfg.DatabasePassword)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DatabaseUsername,
		pass,
		fmt.Sprintf("%s:%s", cfg.DatabaseHost, cfg.DatabasePort),
		cfg.DatabaseName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}
	defer db.Close()

	if _, err := os.Stat(cfg.MigrationPath); os.IsNotExist(err) {
		return fmt.Errorf("migration path does not exist: %s", cfg.MigrationPath)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("Failed to create migration driver: %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", cfg.MigrationPath),
		"postgres", driver)
	if err != nil {
		log.Printf("Failed to initialize migration: %v", err)
		return err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply.")
		} else {
			log.Printf("Migration failed: %v", err)
			return err
		}
	}

	log.Println("Migrations applied successfully.")
	return nil
}
