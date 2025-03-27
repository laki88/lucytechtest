package handlers

import (
	"net/http"
	"regexp"

	"web-analyzer/services"

	"github.com/gin-gonic/gin"
)

type AnalyzeRequest struct {
	URL string `json:"url"`
}

func AnalyzeHandler(c *gin.Context) {
	var req AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	url := req.URL

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}
	if !isValidURL(url) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	go services.AnalyzePage(url)

	c.JSON(http.StatusAccepted, gin.H{"message": "URL submitted for analysis"})
}

func isValidURL(url string) bool {
	// Define a regex pattern for URL validation
	const urlPattern = `^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}(:[0-9]{1,5})?(/.*)?$`

	re := regexp.MustCompile(urlPattern)
	return re.MatchString(url)
}
