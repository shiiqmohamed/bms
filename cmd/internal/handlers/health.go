package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shiiqmohamed/bms/internal/database"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Database  string    `json:"database"`
	Version   string    `json:"version"`
}

// HealthCheck handles health check requests
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Database:  "postgresql",
		Version:   "1.0.0",
	}

	// Check database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.GetDB()
	if db == nil {
		response.Status = "unhealthy"
		response.Database = "disconnected"
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		err := db.Ping(ctx)
		if err != nil {
			response.Status = "unhealthy"
			response.Database = "connection failed"
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ReadyCheck checks if service is ready to accept traffic
func ReadyCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.GetDB()
	if db == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Database not connected"))
		return
	}

	err := db.Ping(ctx)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Database ping failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}