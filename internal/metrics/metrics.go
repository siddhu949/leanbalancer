package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Define Prometheus metrics
var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "leanbalancer_requests_total",
			Help: "Total number of requests processed",
		},
		[]string{"method", "path", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "leanbalancer_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	ActiveConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "leanbalancer_active_connections",
			Help: "Current number of active connections",
		},
	)
)

// Register metrics with Prometheus
func RegisterMetrics() {
	prometheus.MustRegister(RequestsTotal, RequestDuration, ActiveConnections)
}

// Metrics handler for Fasthttp
func MetricsHandler(ctx *fasthttp.RequestCtx) {
	h := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	h(ctx)
}
