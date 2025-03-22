package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"web-analyzer/handlers"
	"web-analyzer/models"
	"web-analyzer/services"
)

func TestStatusHandler_ValidURL(t *testing.T) {
	testURL := "https://example.com"
	services.StoreAnalysis(testURL, models.AnalysisResult{Status: "Completed"})

	req, _ := http.NewRequest("GET", "/status?url="+testURL, nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, status)
	}
}

func TestStatusHandler_MissingURL(t *testing.T) {
	req, _ := http.NewRequest("GET", "/status", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %v", http.StatusBadRequest, status)
	}
}
