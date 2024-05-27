package main

import (
	"betting-backend/api"
	"betting-backend/events-processor"
	"betting-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func initDB() *gorm.DB {
	// Read environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" || dbPassword == "" || dbSSLMode == "" {
		log.Fatalf("Database environment variables are not set properly")
	}

	// Create the connection string
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=" + dbSSLMode

	// Connect to the PostgreSQL database
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the schema (create or update tables based on the model definitions)
	db.AutoMigrate(&model.Bet{}, &model.Event{})

	return db
}

func main() {
	// Initialize the database
	db := initDB()
	defer db.Close()

	// Start the events processor in a separate goroutine
	go events_processor.ReadEvents(db) // Start reading events in a separate goroutine

	// Start the Gin server on the specified port
	r := gin.Default()
	api.InitAPIs(db, r)

	// Start the Gin server on the specified port
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
