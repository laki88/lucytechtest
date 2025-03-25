package handlers

import (
	"net/http"

	"web-analyzer/services"

	"github.com/gin-gonic/gin"
)

func StatusHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}

	analysis, exists := services.GetAnalysis(url)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Analysis not found"})
		return
	}

	c.JSON(http.StatusOK, analysis)
}
