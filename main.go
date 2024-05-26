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

	// Create the connection string
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=" + dbSSLMode

	// Connect to the database
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	db.AutoMigrate(&model.Bet{}, &model.Event{})

	return db
}

func main() {
	// init DB
	db := initDB()
	defer db.Close()

	// start events processor
	go events_processor.ReadEvents(db) // Start reading events in a separate goroutine

	// init APIs and start server
	r := gin.Default()
	api.InitAPIs(db, r)
	err := r.Run(os.Getenv("HTTP_PORT")) // Run on port 8080
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
