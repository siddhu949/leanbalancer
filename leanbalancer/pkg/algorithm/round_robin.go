package algorithm

import (
	"net/url"
	"sync"

	"github.com/siddhu949/leanbalancer/internal/health"
)

// RoundRobin struct
type RoundRobin struct {
	healthChecker *health.HealthChecker
	mu            sync.Mutex
	index         int
}

// NewRoundRobin initializes Round Robin with health check
func NewRoundRobin(hc *health.HealthChecker) *RoundRobin {
	return &RoundRobin{healthChecker: hc}
}

// GetNextBackend selects the next available backend
func (rr *RoundRobin) GetNextBackend() *url.URL {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	healthyBackends := rr.healthChecker.GetHealthyBackends()
	if len(healthyBackends) == 0 {
		return nil // No available servers
	}

	backend, _ := url.Parse(healthyBackends[rr.index])
	rr.index = (rr.index + 1) % len(healthyBackends)
	return backend
}
