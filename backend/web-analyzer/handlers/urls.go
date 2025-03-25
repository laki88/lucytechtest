package handlers

import (
	"net/http"
	"web-analyzer/services"

	"github.com/gin-gonic/gin"
)

func UrlsHandler(c *gin.Context) {
	urls := services.GetSubmittedUrls()

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
