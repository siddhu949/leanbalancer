package pool

import (
	"sync"

	"github.com/valyala/fasthttp"
)

var (
	clientPool = sync.Pool{
		New: func() interface{} {
			return &fasthttp.Client{} // Memory-efficient HTTP client reuse
		},
	}
)

// GetClient retrieves an HTTP client from the pool
func GetClient() *fasthttp.Client {
	return clientPool.Get().(*fasthttp.Client)
}

// ReleaseClient returns an HTTP client to the pool
func ReleaseClient(client *fasthttp.Client) {
	clientPool.Put(client)
}
