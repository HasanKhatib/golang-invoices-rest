package db

import (
	"fmt"
	"log"

	"hasankhatib/golang-invoices-rest/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect connects to the database using the provided configuration
func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	log.Println("Connected to the database")

	return db, nil
}

// Close closes the database connection
func Close(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Failed to get underlying DB connection: %v", err)
			return
		}
		sqlDB.Close()
		log.Println("Database connection closed")
	}
}
