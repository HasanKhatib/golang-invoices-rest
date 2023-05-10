package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
	"hasankhatib/golang-invoices-rest/internal/config"
	"hasankhatib/golang-invoices-rest/internal/db"
	"hasankhatib/golang-invoices-rest/internal/handlers"
	"hasankhatib/golang-invoices-rest/internal/repositories"
	"hasankhatib/golang-invoices-rest/internal/services"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("../configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	// Perform database migration
	if err := runDatabaseMigrations(cfg.Database); err != nil {
		log.Fatalf("Failed to run database migrations: %s", err)
	}

	// Initialize database connection
	dbConn, err := db.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close(dbConn) // Close the database connection

	// Initialize router
	router := mux.NewRouter()

	// Create an instance of the repository
	invoiceRepository := repositories.NewInvoiceRepository(dbConn)

	// Create an instance of the service
	invoiceService := services.NewInvoiceService(invoiceRepository)

	// Initialize invoice handler
	invoiceHandler := handlers.NewInvoiceHandler(*invoiceService)

	// Define routes
	router.HandleFunc("/invoices", invoiceHandler.GetAllInvoices).Methods("GET")
	router.HandleFunc("/invoices/{id}", invoiceHandler.GetInvoiceByID).Methods("GET")
	router.HandleFunc("/invoices", invoiceHandler.CreateInvoice).Methods("POST")

	// Start server
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server started at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func runDatabaseMigrations(dbCfg config.DatabaseConfig) error {
	// Create database connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.Name, dbCfg.Port)

	// Open connection to the database
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Create a new PostgreSQL driver instance
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create PostgreSQL driver instance: %v", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply database migrations: %v", err)
	}

	return nil
}
