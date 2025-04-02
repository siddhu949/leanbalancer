package health

import (
	"net/http"
	"sync"
	"time"
)

// Backend represents a backend server
type Backend struct {
	URL   string
	Alive bool
	mu    sync.RWMutex
}

// HealthChecker maintains backend health status
type HealthChecker struct {
	Backends []*Backend
	Interval time.Duration
}

// NewHealthChecker initializes the health checker
func NewHealthChecker(backendURLs []string, interval time.Duration) *HealthChecker {
	backends := make([]*Backend, len(backendURLs))
	for i, url := range backendURLs {
		backends[i] = &Backend{URL: url, Alive: true}
	}

	return &HealthChecker{Backends: backends, Interval: interval}
}

// CheckHealth verifies backend status and updates their availability
func (hc *HealthChecker) CheckHealth() {
	for {
		var wg sync.WaitGroup
		for _, backend := range hc.Backends {
			wg.Add(1)
			go func(b *Backend) {
				defer wg.Done()
				client := &http.Client{Timeout: 2 * time.Second}
				resp, err := client.Get(b.URL + "/health") // Expecting /health endpoint
				b.mu.Lock()
				if err == nil && resp.StatusCode == http.StatusOK {
					b.Alive = true
				} else {
					b.Alive = false
				}
				b.mu.Unlock()
			}(backend)
		}
		wg.Wait()
		time.Sleep(hc.Interval)
	}
}

// GetHealthyBackends returns a list of available backends
func (hc *HealthChecker) GetHealthyBackends() []string {
	healthy := []string{}
	for _, backend := range hc.Backends {
		backend.mu.RLock()
		if backend.Alive {
			healthy = append(healthy, backend.URL)
		}
		backend.mu.RUnlock()
	}
	return healthy
}
