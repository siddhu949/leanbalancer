package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus Metrics
var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_requests_total",
			Help: "Total number of requests received by backend",
		},
		[]string{"backend_id", "path"},
	)
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🩺 Received health check request")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// Generic handler for both "/" and "/reverse"
func handleRequest(backendID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		requestsTotal.WithLabelValues(backendID, path).Inc()
		log.Printf("✅ Received request at %s on Backend %s\n", path, backendID)
		fmt.Fprintf(w, "✅ Response from %s on Backend %s\n", path, backendID)
	}
}

func main() {
	// Read Backend ID and Port from Env Variables
	backendID := os.Getenv("BACKEND_ID")
	if backendID == "" {
		backendID = "1"
	}

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "9001"
	}

	// Register Prometheus metrics
	prometheus.MustRegister(requestsTotal)

	// Create HTTP Server
	server := &http.Server{
		Addr: ":" + port,
	}

	// Register Handlers
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", handleRequest(backendID))
	http.HandleFunc("/reverse", handleRequest(backendID)) // ✅ New route

	// Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	// Graceful shutdown setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server
	go func() {
		log.Printf("🚀 Backend %s running on port: %s", backendID, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Backend %s failed to start: %v", backendID, err)
		}
	}()

	<-stop
	log.Println("🔄 Shutting down backend gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Backend shutdown error: %v", err)
	} else {
		log.Println("✅ Backend stopped successfully")
	}
}
