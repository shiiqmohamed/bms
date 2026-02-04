package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/shiiqmohamed/bms/internal/config"
	"github.com/shiiqmohamed/bms/internal/database"
	"github.com/shiiqmohamed/bms/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()
	
	fmt.Println("🚀 Starting BMS API...")
	
	// Initialize database
	if err := database.InitializeDB(); err != nil {
		log.Fatal("Database error:", err)
	}
	defer database.CloseDB()
	
	// Setup routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/health", handlers.HealthCheck)
	
	// Start server
	port := cfg.ServerPort
	log.Printf("✅ Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
