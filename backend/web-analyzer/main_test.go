package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMainServer(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode

	// Start the server in a test
	go func() {
		main()
	}()
	time.Sleep(2 * time.Second) // Give server time to start

	// Test if the `/metrics` endpoint is available
	resp, err := http.Get("http://localhost:8080/metrics")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAnalyzeEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter() // Use extracted function to create router

	req, _ := http.NewRequest("POST", "/analyze", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting 400 because of missing payload
}
