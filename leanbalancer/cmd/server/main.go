package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "github.com/siddhu949/leanbalancer/api/v1"
	"github.com/siddhu949/leanbalancer/internal/admin"
	"github.com/siddhu949/leanbalancer/internal/firewall"
	"github.com/siddhu949/leanbalancer/internal/health"
	"github.com/siddhu949/leanbalancer/internal/logger"
	"github.com/siddhu949/leanbalancer/internal/metrics"
	"github.com/siddhu949/leanbalancer/internal/proxy"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// Define configuration values directly here
var serverPort = 8080
var metricsPort = 9090
var loadBalancerAlgorithm = "round_robin"
var timeout = 5 * time.Second

var firewallEnabled = true
var blockedIPs = []string{"192.168.1.100", "10.0.0.1"}

var healthCheckEnabled = true
var healthCheckInterval = 5 * time.Second
var healthCheckBackends = []string{
	"http://localhost:9001",
	"http://localhost:9002",
	"http://localhost:9003",
}

// Request Router to handle different requests
func requestHandler(ctx *fasthttp.RequestCtx) {
	// Apply Firewall Check Before Proxy
	if firewallEnabled && !firewall.FirewallMiddleware(ctx) {
		return
	}

	// Routing for Reverse Proxy, Forward Proxy and Health
	switch string(ctx.Path()) {
	case "/reverse":
		proxy.ReverseProxyHandler(ctx)
	case "/forward":
		proxy.ForwardProxyHandler(ctx)
	case "/health":
		ctx.SetStatusCode(200)
		ctx.SetBody([]byte("OK"))
	case "/metrics":
		// Handle Prometheus metrics scraping
		handleMetrics(ctx)
	default:
		ctx.Error("Not Found", fasthttp.StatusNotFound)
	}
}

// Prometheus metrics handler for fasthttp
func handleMetrics(ctx *fasthttp.RequestCtx) {
	// Create a wrapped response writer to satisfy http.ResponseWriter interface
	wrappedWriter := &responseWriterWrapper{ctx}

	// Serve the Prometheus metrics
	promhttp.Handler().ServeHTTP(wrappedWriter, &http.Request{})
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Log health check request
	log.Println("Received health check request")

	// Respond with a 200 OK status and "OK" message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// Graceful shutdown handler
func gracefulShutdown(srv *http.Server) {
	// Wait for an interrupt signal (Ctrl+C) to shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutting down gracefully...")

	// Give current requests 5 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server stopped gracefully.")
}

// ResponseWriterWrapper is used to wrap fasthttp's RequestCtx to fit http.ResponseWriter
type responseWriterWrapper struct {
	*fasthttp.RequestCtx
}

func (w *responseWriterWrapper) Header() http.Header {
	// Returning an empty header as we don't need to interact with headers
	return http.Header{}
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	// Set the status code in the fasthttp context
	w.SetStatusCode(statusCode)
}

func (w *responseWriterWrapper) Write(p []byte) (n int, err error) {
	// Write the data to the fasthttp context response body
	w.SetBody(p)
	return len(p), nil
}

func main() {
	// Initialize Logger
	logger.InitLogger()
	log := logger.GetLogger()

	// Start Fiber app
	app := fiber.New()
	v1.SetupRoutes(app)            // Correctly setup routes for API v1
	admin.RegisterAdminRoutes(app) // Register admin routes

	// Initialize Health Checker if enabled
	if healthCheckEnabled {
		healthChecker := health.NewHealthChecker(healthCheckBackends, healthCheckInterval)
		go healthChecker.CheckHealth()
		log.Info("Health check enabled", zap.Int("backends", len(healthCheckBackends)))
	}

	// Register Prometheus metrics
	metrics.RegisterMetrics()

	// Create the server instance
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: http.DefaultServeMux,
	}

	// Start API Server in a separate goroutine
	go func() {
		log.Info("✅ API v1 running on port", zap.Int("port", metricsPort))
		if err := app.Listen(":9092"); err != nil {
			log.Fatal("Error starting API server", zap.Error(err))
		}
	}()

	// Start Load Balancer Server
	log.Info("LeanBalancer running on port", zap.Int("port", serverPort))
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", serverPort), requestHandler); err != nil {
		log.Fatal("Error starting server", zap.Error(err))
	}

	// Graceful shutdown handler for the HTTP server
	gracefulShutdown(srv)
}
