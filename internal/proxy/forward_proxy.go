package proxy

import (
	"fmt"
	"time"

	"github.com/siddhu949/leanbalancer/pkg/pool"
	"github.com/siddhu949/leanbalancer/pkg/utils"
	"github.com/valyala/fasthttp"
)

// Forward Proxy Handler
func ForwardProxyHandler(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()

	target := string(ctx.QueryArgs().Peek("target"))
	if target == "" {
		ctx.Error("Missing target parameter", fasthttp.StatusBadRequest)
		return
	}

	client := pool.GetClient()
	defer pool.ReleaseClient(client)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	ctx.Request.CopyTo(req)
	req.SetRequestURI(target)

	err := client.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		ctx.Error(fmt.Sprintf("Error forwarding request: %s", err), fasthttp.StatusServiceUnavailable)
		return
	}

	utils.LogRequest(string(ctx.Method()), target, resp.StatusCode(), time.Since(startTime))

	resp.CopyTo(&ctx.Response)
}
