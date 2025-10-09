package initiator

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yohannesgossaye/internal/persistence/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func InitDatabase() *pgxpool.Pool {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get database configuration from environment
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "users_db")

	// Create database connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Create connection pool
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parse database config: %v", err)
	}

	// Set connection pool settings
	config.MaxConns = 10
	config.MinConns = 2

	// Create connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to create database connection pool: %v", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Run migrations
	if err := postgres.RunMigrations(connStr); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connection established successfully")
	return pool
}

func InitRepository(db *pgxpool.Pool) *postgres.CustomerRepositary {
	repo := postgres.NewCustomerRepositary(db)
	return repo
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
