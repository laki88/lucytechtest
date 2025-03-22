package main

import (
	"log"
	"net/http"

	"web-analyzer/handlers"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default() // Ensures logging and recovery middleware is enabled

	// Define API routes
	r.POST("/analyze", func(c *gin.Context) {
		handlers.AnalyzeHandler(c.Writer, c.Request)
	})
	r.GET("/status", func(c *gin.Context) {
		handlers.StatusHandler(c.Writer, c.Request)
	})
	r.GET("/urls", func(c *gin.Context) {
		handlers.UrlsHandler(c.Writer, c.Request)
	})

	// Enable CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust based on frontend
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(r)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
