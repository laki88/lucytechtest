package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"web-analyzer/handlers"
)

func TestAnalyzeHandler_ValidURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/analyze", strings.NewReader("url=https://example.com"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AnalyzeHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("Expected status %v, got %v", http.StatusAccepted, status)
	}
}

func TestAnalyzeHandler_MissingURL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/analyze", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AnalyzeHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %v", http.StatusBadRequest, status)
	}
}
