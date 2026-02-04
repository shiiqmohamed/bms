package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shiiqmohamed/bms/internal/config"
)

var DB *pgxpool.Pool

// InitializeDB initializes the database connection pool
func InitializeDB() error {
	cfg := config.LoadConfig()
	
	connString := cfg.GetDBConnectionString()
	
	// Configure connection pool
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return fmt.Errorf("unable to parse database config: %w", err)
	}

	// Set connection pool parameters
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = time.Minute * 30
	poolConfig.HealthCheckPeriod = time.Minute

	// Create connection pool
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DB, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Test the connection
	if err := testConnection(ctx); err != nil {
		return fmt.Errorf("database connection test failed: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

// testConnection verifies the database connection
func testConnection(ctx context.Context) error {
	conn, err := DB.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	return conn.Ping(ctx)
}

// CloseDB closes the database connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}

// GetDB returns the database connection pool
func GetDB() *pgxpool.Pool {
	return DB
}