package database

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/YusronAbi/ecomerce-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(ctx context.Context, cfg config.Config) *gorm.DB {
	pass := url.QueryEscape(cfg.DatabasePassword)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DatabaseUsername,
		pass,
		fmt.Sprintf("%s:%s", cfg.DatabaseHost, cfg.DatabasePort),
		cfg.DatabaseName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Printf("Connected to the database")

	return db
}