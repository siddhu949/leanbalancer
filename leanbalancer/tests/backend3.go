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
		[]string{"backend_id"},
	)
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Log health check request
	log.Println("Received health check request")

	// Respond with a 200 OK status and "OK" message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	// Read Backend ID from Env Variable or Default to 1
	backendID := os.Getenv("BACKEND_ID")
	if backendID == "" {
		backendID = "3"
	}

	// Read Port from Env Variable or Default to 9001
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "9003"
	}

	// Register Prometheus metrics
	prometheus.MustRegister(requestsTotal)

	// Create HTTP Server
	server := &http.Server{
		Addr: ":" + port,
	}

	// Register health check handler
	http.HandleFunc("/health", healthHandler)

	// HTTP Handler for Backend Response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment Prometheus counter
		requestsTotal.WithLabelValues(backendID).Inc()

		// Log request and respond
		log.Println("✅ Received request at Backend", backendID)
		fmt.Fprintf(w, "✅ Response from Backend %s\n", backendID)
	})

	// Expose Prometheus Metrics at /metrics
	http.Handle("/metrics", promhttp.Handler())

	// Graceful Shutdown Setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		log.Printf("🚀 Backend %s running on port: %s", backendID, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Backend %s failed to start: %v", backendID, err)
		}
	}()

	// Wait for termination signal
	<-stop
	log.Println("🔄 Shutting down backend gracefully...")

	// Graceful Shutdown Logic
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Backend shutdown error: %v", err)
	} else {
		log.Println("✅ Backend stopped successfully")
	}
}
