package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shiiqmohamed/bms/internal/database"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Welcome to BMS API",
		"status":  "OK",
		"time":    time.Now().Format(time.RFC3339),
	})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"service":   "BMS API",
		"database":  "connected",
	}

	db := database.GetDB()
	if db == nil {
		response["status"] = "unhealthy"
		response["database"] = "disconnected"
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
