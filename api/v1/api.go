package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartAPIServer() {
	app := fiber.New()

	// Set up routes from routes.go
	SetupRoutes(app)

	// Start the server
	log.Println("✅ API v1 running on port 9090")
	go func() {
		if err := app.Listen(":9090"); err != nil {
			log.Fatalf("❌ Failed to start API server: %v", err)
		}
	}()
}
