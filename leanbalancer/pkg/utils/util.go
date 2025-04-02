package utils

import (
	"log"
	"time"
)

// LogRequest logs details about each request
func LogRequest(method, path string, status int, duration time.Duration) {
	log.Printf("[%s] %s -> %d (%v)\n", method, path, status, duration)
}
