package services_test

import (
	"testing"
	"web-analyzer/services"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzePage_ValidURL(t *testing.T) {
	url := "http://example.com"
	services.AnalyzePage(url)
	result, exists := services.GetAnalysis(url)

	assert.True(t, exists, "Analysis result should exist")
	assert.NotEmpty(t, result.Status, "Status should not be empty")
}

func TestAnalyzePage_InvalidURL(t *testing.T) {
	url := "invalid-url"
	services.AnalyzePage(url)
	result, exists := services.GetAnalysis(url)

	assert.True(t, exists, "Even invalid URLs should be stored")
	assert.Equal(t, "Error", result.Status, "Status should be 'Error'")
}
