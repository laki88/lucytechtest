package services

import (
	"testing"

	"web-analyzer/models"
	"web-analyzer/services"
)

func TestStoreAnalysis(t *testing.T) {
	url := "https://example.com"
	expected := models.AnalysisResult{Status: "Completed"}

	services.StoreAnalysis(url, expected)
	result, exists := services.GetAnalysis(url)

	if !exists || result.Status != expected.Status {
		t.Errorf("Expected status %v, got %v", expected.Status, result.Status)
	}
}

func TestGetAnalysis_NotFound(t *testing.T) {
	_, exists := services.GetAnalysis("https://nonexistent.com")

	if exists {
		t.Error("Expected non-existent analysis to return false")
	}
}
