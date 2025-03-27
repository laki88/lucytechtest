package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"web-analyzer/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzeHandler_ValidURL(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/analyze", handlers.AnalyzeHandler)

	body, _ := json.Marshal(map[string]string{"url": "http://example.com"})
	req, _ := http.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusAccepted, resp.Code)
}

func TestAnalyzeHandler_InvalidURL(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/analyze", handlers.AnalyzeHandler)

	body, _ := json.Marshal(map[string]string{"url": "invalid-url"})
	req, _ := http.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
