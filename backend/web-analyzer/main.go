package main

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"web-analyzer/handlers"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.Default() // Ensures logging and recovery middleware is enabled

	setupPprof(r)

	pprof.Register(r)

	// Prometheus Metrics
	prometheusRegistry := prometheus.NewRegistry()
	httpRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
	prometheusRegistry.MustRegister(httpRequests)

	r.Use(func(c *gin.Context) {
		httpRequests.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		c.Next()
	})

	// Expose Prometheus metrics at /metrics
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})))

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

	logger.Info("Server started on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		logger.Error("Server failed to start", "error", err)
	}
}

func setupPprof(router *gin.Engine) {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
}
