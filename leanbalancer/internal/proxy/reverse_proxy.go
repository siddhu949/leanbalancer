package proxy

import (
	"fmt"
	"time"

	"github.com/siddhu949/leanbalancer/internal/health"
	"github.com/siddhu949/leanbalancer/pkg/algorithm"
	"github.com/siddhu949/leanbalancer/pkg/pool"
	"github.com/siddhu949/leanbalancer/pkg/utils"
	"github.com/valyala/fasthttp"
)

// Initialize Health Check
var healthChecker = health.NewHealthChecker([]string{
	"http://localhost:9001",
	"http://localhost:9002",
	"http://localhost:9003",
}, 5*time.Second)

// Initialize Round Robin with Health Check
var roundRobin = algorithm.NewRoundRobin(healthChecker)

// Reverse Proxy Handler
func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()
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

	ctx.Request.CopyTo(req)
	req.SetRequestURI(backend.String() + string(ctx.Path()))

	err := client.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		ctx.Error(fmt.Sprintf("Error connecting to backend: %s", err), fasthttp.StatusServiceUnavailable)
		return
	}

	utils.LogRequest(string(ctx.Method()), string(ctx.Path()), resp.StatusCode(), time.Since(startTime))

	resp.CopyTo(&ctx.Response)
}
