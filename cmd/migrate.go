package main

import (
	"log"

	"hasanalkhatib/rest-service/internal/config"
	"hasanalkhatib/rest-service/internal/db"
	"hasanalkhatib/rest-service/internal/models"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Connect to the database
	err = db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Auto-migrate the Invoice model
	if err := db.DB.AutoMigrate(&models.Invoice{}); err != nil {
		log.Fatalf("failed to migrate the database: %v", err)
	}

	log.Println("Database migration complete.")
}
