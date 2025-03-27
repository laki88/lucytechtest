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
	slog.SetDefault(logger)

	r := setupRouter()

	pprof.Register(r)

	logger.Info("Setting up Prometheus metrics")
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
		logger.Info("Received request", "method", c.Request.Method, "endpoint", c.FullPath())
		c.Next()
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
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			slog.Error("pprof server failed", "error", err)
		}
	}()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	setupPprof(r)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.POST("/analyze", handlers.AnalyzeHandler)
	r.GET("/status", handlers.StatusHandler)
	r.GET("/urls", handlers.UrlsHandler)

	return r
}
