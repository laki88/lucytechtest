package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web-analyzer/handlers"
	"web-analyzer/services"
)

func TestUrlsHandler(t *testing.T) {
	services.AddSubmittedUrl("https://example.com")

	req, _ := http.NewRequest("GET", "/urls", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.UrlsHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}
}
