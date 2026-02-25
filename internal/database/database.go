package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB() error {
	// Use environment variables directly
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Debug logging
	log.Printf("Attempting to connect to database: %s@%s:%s/%s", user, host, port, dbname)

	// IMPORTANT: sslmode=require for Neon.tech
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require connect_timeout=10",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Try to ping with retry
	for i := 0; i < 3; i++ {
		err = DB.Ping()
		if err == nil {
			break
		}
		log.Printf("Ping attempt %d failed: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to ping database after retries: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("📴 Database connection closed")
	}
}

func GetDB() *sql.DB {
	return DB
}
