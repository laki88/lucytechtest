package handlers

import (
	"encoding/json"
	"net/http"

	"web-analyzer/services"
)

func UrlsHandler(w http.ResponseWriter, r *http.Request) {
	urls := services.GetSubmittedUrls()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"urls": urls})
}
