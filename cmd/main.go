package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shiiqmohamed/bms/internal/database"
	"github.com/shiiqmohamed/bms/internal/handlers"
)

func main() {
	// Debug: Print environment variables
	log.Println("🔍 Environment variables:")
	log.Printf("   DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("   DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("   DB_NAME: %s", os.Getenv("DB_NAME"))
	log.Printf("   DB_PORT: %s", os.Getenv("DB_PORT"))
	log.Printf("   SERVER_PORT: %s", os.Getenv("SERVER_PORT"))

	log.Println("🚀 Starting BMS API...")

	err := database.InitializeDB()
	if err != nil {
		log.Fatalf("❌ Database error:%v", err)
	}
	defer database.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheck)

	server := &http.Server{
		Addr:         "0.0.0.0:" + os.Getenv("SERVER_PORT"),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("✅ Server running on http://localhost:%s", os.Getenv("SERVER_PORT"))

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("❌ server failed: %v", err)
	}
}
