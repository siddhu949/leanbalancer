package proxy

import (
	"fmt"
	"strings"
	"time"

	"github.com/siddhu949/leanbalancer/internal/health"
	"github.com/siddhu949/leanbalancer/pkg/algorithm"
	"github.com/siddhu949/leanbalancer/pkg/pool"
	"github.com/siddhu949/leanbalancer/pkg/utils"
	"github.com/valyala/fasthttp"
)

// Health Checker initialized with backends
var healthChecker = health.NewHealthChecker([]string{
	"http://localhost:9001",
	"http://localhost:9002",
	"http://localhost:9003",
}, 5*time.Second)

// Round Robin instance
var roundRobin = algorithm.NewRoundRobin(healthChecker)

// ReverseProxyHandler handles reverse proxy requests
func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	start := time.Now()
	backend := roundRobin.GetNextBackend()

	if backend == nil {
		ctx.Error("No available backends", fasthttp.StatusServiceUnavailable)
		return
	}

	client := pool.GetClient()
	defer pool.ReleaseClient(client)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// Copy incoming request
	ctx.Request.CopyTo(req)

	// Clean path (remove /reverse)
	cleanPath := strings.TrimPrefix(string(ctx.Path()), "/reverse")

	// Rebuild the new URI
	req.SetRequestURI(backend.String() + cleanPath)

	// Set the correct Host header for the backend
	req.SetHost(backend.Host)

	// Optional: copy headers (already done via CopyTo, but you can double-check)
	// ctx.Request.Header.CopyTo(&req.Header)

	// Perform request with timeout
	err := client.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		ctx.Error(fmt.Sprintf("Error forwarding request: %s", err), fasthttp.StatusServiceUnavailable)
		return
	}

	// Log and send response
	utils.LogRequest(string(ctx.Method()), string(ctx.Path()), resp.StatusCode(), time.Since(start))
	resp.CopyTo(&ctx.Response)
}
