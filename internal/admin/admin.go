package admin

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
)

// AdminHandler serves Prometheus metrics over an admin endpoint
func AdminHandler(ctx *fasthttp.RequestCtx) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprint(ctx, "Internal Server Error")
		return
	}

	// Use httptest to record the response
	recorder := httptest.NewRecorder()

	// Serve Prometheus metrics using http
	metricsHandler := promhttp.Handler()
	metricsHandler.ServeHTTP(recorder, req)

	// Copy the recorded response back to fasthttp
	for k, v := range recorder.Header() {
		ctx.Response.Header.Set(k, v[0])
	}
	ctx.SetStatusCode(recorder.Code)
	ctx.Write(recorder.Body.Bytes()) // Send response
}

// Health check endpoint
func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "running",
	})
}

// RegisterAdminRoutes registers admin routes
func RegisterAdminRoutes(app *fiber.App) {
	app.Get("/admin", func(c *fiber.Ctx) error {
		return c.SendString("Admin Panel")
	})
}

// Example handler with Fiber header fix
func ExampleHandler(ctx *fiber.Ctx) error {
	contentType := ctx.Get(fiber.HeaderContentType) // Correct usage
	log.Println("Received Content-Type:", contentType)

	ctx.Response().Header.Set("Content-Type", "application/json") // Set response header
	return ctx.JSON(fiber.Map{
		"message": "Hello from LeanBalancer",
	})
}
