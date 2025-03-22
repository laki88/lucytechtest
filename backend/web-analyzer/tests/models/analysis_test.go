package models

import (
	"strings"
	"testing"

	"web-analyzer/models"

	"golang.org/x/net/html"
)

func TestAnalyzeHTML(t *testing.T) {
	rawHTML := `<html><head><title>Test Page</title></head><body><h1>Header</h1><a href="https://example.com"></a><form action="login"></form></body></html>`
	doc, _ := html.Parse(strings.NewReader(rawHTML))

	result := models.AnalyzeHTML(doc, "https://test.com")

	if result.Title != "Test Page" {
		t.Errorf("Expected title 'Test Page', got %s", result.Title)
	}
	if result.Headings["h1"] != 1 {
		t.Errorf("Expected 1 h1 tag, got %d", result.Headings["h1"])
	}
	if !result.LoginForm {
		t.Errorf("Expected login form detection")
	}
}
