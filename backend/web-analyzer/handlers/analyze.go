package handlers

import (
	"encoding/json"
	"net/http"

	"web-analyzer/services"
)

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	go services.AnalyzePage(url)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "URL submitted for analysis"})
}
