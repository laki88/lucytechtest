package models_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"web-analyzer/models"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestDetectHTMLVersion_HTML5(t *testing.T) {
	htmlContent := `<!DOCTYPE html><html><head></head><body></body></html>`
	doc, _ := html.Parse(strings.NewReader(htmlContent))
	version := models.DetectHTMLVersion(doc)
	assert.Equal(t, "HTML5", version)
}

func TestDetectHTMLVersion_Unknown(t *testing.T) {
	htmlContent := `<html><head></head><body></body></html>`
	doc, _ := html.Parse(strings.NewReader(htmlContent))
	version := models.DetectHTMLVersion(doc)
	assert.Equal(t, "No DOCTYPE found", version)
}

func TestAnalyzeHTML_ValidHTML(t *testing.T) {
	htmlContent := `<!DOCTYPE html><html><head><title>Test</title></head><body>
		<h1>Heading</h1>
		<a href="http://example.com">External Link</a>
		<a href="/internal">Internal Link</a>
	</body></html>`
	doc, _ := html.Parse(strings.NewReader(htmlContent))
	result := models.AnalyzeHTML(doc, "http://example.com")

	assert.Equal(t, "Test", result.Title)
	assert.Equal(t, 1, result.Headings["h1"])
	assert.Equal(t, 1, result.ExternalLinks)
	assert.Equal(t, 1, result.InternalLinks)
}

func TestIsBrokenLink_FakeURL(t *testing.T) {
	broken := models.IsBrokenLink("http://invalid.url")
	assert.True(t, broken, "Fake URL should be considered broken")
}

func TestIsBrokenLink_ValidURL(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close() // Ensure server shuts down after test

	broken := models.IsBrokenLink(server.URL)
	assert.False(t, broken, "Test server should not be considered broken")
}
