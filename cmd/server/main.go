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
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

var (
	serverPort            = 8080
	metricsPort           = 9090
	loadBalancerAlgorithm = "round_robin"
	timeout               = 5 * time.Second

	firewallEnabled = true
	//blockedIPs      = []string{"192.168.1.100", "10.0.0.1"}

	healthCheckEnabled  = true
	healthCheckInterval = 5 * time.Second
	healthCheckBackends = []string{
		"http://localhost:9001",
		"http://localhost:9002",
		"http://localhost:9003",
	}
)

// LeanBalancer handler
func requestHandler(ctx *fasthttp.RequestCtx) {
	defer func() {
		log.Printf("Responded with status: %d to path: %s", ctx.Response.StatusCode(), ctx.Path())
	}()

	// Firewall check
	if firewallEnabled && !firewall.FirewallMiddleware(ctx) {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		ctx.SetContentType("text/plain")
		ctx.SetBody([]byte("Access denied"))
		return
	}

	switch string(ctx.Path()) {
	case "/":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("text/plain")
		ctx.SetBody([]byte("LeanBalancer OK"))

	case "/reverse":
		proxy.ReverseProxyHandler(ctx)

	case "/forward":
		proxy.ForwardProxyHandler(ctx)

	case "/health":
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("text/plain")
		ctx.SetBody([]byte("OK"))

	case "/metrics":
		fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())(ctx)

	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("text/plain")
		ctx.SetBody([]byte("404 - Not Found"))
	}
}

// Graceful shutdown logic
func gracefulShutdown(srv *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server stopped.")
}

func main() {
	// Init logger
	logger.InitLogger()
	log := logger.GetLogger()

	// Fiber for Admin/API
	app := fiber.New()
	v1.SetupRoutes(app)
	admin.RegisterAdminRoutes(app)

	// Health Check Setup
	if healthCheckEnabled {
		healthChecker := health.NewHealthChecker(healthCheckBackends, healthCheckInterval)
		go healthChecker.CheckHealth()
		log.Info("Health checks enabled", zap.Int("backends", len(healthCheckBackends)))
	}

	// Register Prometheus metrics
	metrics.RegisterMetrics()

	// HTTP server for graceful shutdown (used by Fiber)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: http.DefaultServeMux,
	}

	// Start Fiber API
	go func() {
		log.Info("✅ Admin API running on port", zap.Int("port", metricsPort))
		if err := app.Listen(fmt.Sprintf(":%d", metricsPort)); err != nil {
			log.Fatal("Error starting API server", zap.Error(err))
		}
	}()

	// Start main LeanBalancer proxy server
	log.Info("🚀 LeanBalancer running on port", zap.Int("port", serverPort))
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", serverPort), requestHandler); err != nil {
		log.Fatal("Error starting LeanBalancer", zap.Error(err))
	}

	// Handle graceful shutdown
	gracefulShutdown(srv)
}
