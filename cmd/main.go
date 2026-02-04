
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/shiiqmohamed/bms/internal/config"
	"github.com/shiiqmohamed/bms/internal/database"
	"github.com/shiiqmohamed/bms/internal/handlers"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	
	fmt.Println("🚀 Starting BMS (Business Management System)...")
	fmt.Printf("📊 Database: %s\n", cfg.DBName)
	fmt.Printf("🌐 Server Port: %s\n", cfg.ServerPort)

	// Initialize database
	log.Println("🔌 Connecting to database...")
	if err := database.InitializeDB(); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.CloseDB()

	// Setup HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", handlers.HealthCheck)
	http.HandleFunc("/ready", handlers.ReadyCheck)
	http.HandleFunc("/api/v1/status", apiStatusHandler)

	// Start server in goroutine
	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("✅ Server is running on http://localhost:%s\n", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("🛑 Shutting down server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}
	
	log.Println("✅ Server shutdown completed")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Welcome to BMS (Business Management System) API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"health":    "/health",
			"ready":     "/ready",
			"api_status": "/api/v1/status",
		},
		"database": "PostgreSQL",
	}
	json.NewEncoder(w).Encode(response)
}

func apiStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":  "active",
		"service": "bms-api",
		"time":    time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}