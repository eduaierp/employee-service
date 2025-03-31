package main

import (
	"employee-service/internal/db"
	"employee-service/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database connection string from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}

	// Initialize database connection
	err = db.InitDB(dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Start gRPC server
	server.StartServer()
}
