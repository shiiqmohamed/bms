package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shiiqmohamed/bms/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB() error {
	cfg := config.LoadConfig()
	
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
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
