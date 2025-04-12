package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/asterixcode/payment-gateway/go-gateway-api/internal/repository"
	"github.com/asterixcode/payment-gateway/go-gateway-api/internal/service"
	"github.com/asterixcode/payment-gateway/go-gateway-api/internal/web/server"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// db string connection
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("PORT", "8080")
	srv := server.NewServer(accountService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
