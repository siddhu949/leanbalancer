package firewall

import (
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

// Rate limit settings
const (
	requestLimit  = 100             // Max requests per IP
	windowSize    = time.Minute     // Time window for rate limiting
	blockDuration = 5 * time.Minute // Block duration after exceeding limit
)

// Memory-efficient sync.Map for tracking requests & blocked IPs
var (
	requestCounts sync.Map // Stores request counts per IP
	blockedIPs    sync.Map // Stores blocked IPs & unblock time
)

// FirewallMiddleware checks rate limits and blocks IPs if needed
func FirewallMiddleware(ctx *fasthttp.RequestCtx) bool {
	clientIP := ctx.RemoteIP().String()

	// ✅ Check if the IP is blocked
	if unblockTime, exists := getBlockedIP(clientIP); exists && time.Now().Before(unblockTime) {
		ctx.Error("Access Denied: Too many requests", fasthttp.StatusForbidden)
		return false
	}

	// ✅ Update request count
	count := incrementRequestCount(clientIP)
	if count > requestLimit {
		blockIP(clientIP)
		ctx.Error("Too many requests, access blocked", fasthttp.StatusTooManyRequests)
		return false
	}

	return true
}

// Increments request count for an IP, deletes expired entries
func incrementRequestCount(ip string) int {
	// Fetch and increment the request count for this IP
	count, _ := requestCounts.LoadOrStore(ip, 0) // Initialize at 0 if not found
	newCount := count.(int) + 1
	requestCounts.Store(ip, newCount)

	// ✅ Auto-delete after windowSize (Prevents Memory Leak)
	go func() {
		time.Sleep(windowSize)
		requestCounts.Delete(ip) // Free memory after time window
	}()

	return newCount
}

// Blocks an IP for a set duration
func blockIP(ip string) {
	blockedIPs.Store(ip, time.Now().Add(blockDuration))

	// ✅ Auto-remove block after duration
	go func() {
		time.Sleep(blockDuration)
		blockedIPs.Delete(ip) // Free memory after blocking expires
	}()
}

// Gets blocked IP and its unblock time
func getBlockedIP(ip string) (time.Time, bool) {
	value, exists := blockedIPs.Load(ip)
	if exists {
		return value.(time.Time), true
	}
	return time.Time{}, false
}

// GetBlockedIPs returns a list of all blocked IPs
func GetBlockedIPs() []string {
	var blocked []string
	blockedIPs.Range(func(key, value interface{}) bool {
		blocked = append(blocked, key.(string))
		return true
	})
	return blocked
}

// BlockIP adds an IP to the blocked list
func BlockIP(ip string) {
	blockedIPs.Store(ip, time.Now().Add(blockDuration))

	// ✅ Auto-remove block after duration
	go func() {
		time.Sleep(blockDuration)
		blockedIPs.Delete(ip) // Free memory after blocking expires
	}()
}
