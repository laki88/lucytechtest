package handlers

import (
	"encoding/json"
	"net/http"

	"web-analyzer/services"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	analysis, exists := services.GetAnalysis(url)
	if !exists {
		http.Error(w, "Analysis not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}
