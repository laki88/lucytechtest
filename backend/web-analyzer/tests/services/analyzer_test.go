package services

import (
	"testing"

	"web-analyzer/services"
)

func TestAnalyzePage_ValidURL(t *testing.T) {
	testURL := "https://www.example.com"

	go services.AnalyzePage(testURL)
	// No direct assertion as the analysis runs asynchronously.
}

func TestAnalyzePage_InvalidURL(t *testing.T) {
	testURL := "invalid-url"

	go services.AnalyzePage(testURL)
	// Ensure it doesn't panic; check logs manually.
}
