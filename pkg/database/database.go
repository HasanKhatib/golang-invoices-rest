package database

import (
	"fmt"
	"log"

	"hasankhatib/golang-invoices-rest/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB represents the database connection
var DB *gorm.DB

// Connect establishes a connection to the database
func Connect(databaseURL string) error {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Migrate the database schema
	if err := db.AutoMigrate(&models.Invoice{}); err != nil {
		return fmt.Errorf("failed to migrate the database: %v", err)
	}

	DB = db
	return nil
}

// Close closes the database connection
func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("failed to close the database connection: %v", err)
			return
		}
		sqlDB.Close()
	}
}
