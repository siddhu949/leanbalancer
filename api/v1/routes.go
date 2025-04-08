// routes.go
package v1

import (
	"github.com/gofiber/fiber/v2"
)

// Temporary in-memory backend store (for demo purpose)
var backendList = []string{}

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1") // Base route

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "API v1 Running"})
	})

	v1.Get("/health", HealthCheckHandler)
	v1.Get("/firewall", GetFirewallRules)
	v1.Post("/firewall/block", BlockIP)

	// ✅ Add these backend routes
	v1.Post("/backends", AddBackend)
	v1.Get("/backends", GetBackends)
	v1.Delete("/backends/:id", RemoveBackend) // optional, for dynamic removal
}

// ✅ Handler to register a backend
func AddBackend(c *fiber.Ctx) error {
	type Backend struct {
		URL string `json:"url"`
	}
	var backend Backend
	if err := c.BodyParser(&backend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	backendList = append(backendList, backend.URL)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Backend added", "url": backend.URL})
}

// ✅ Handler to list registered backends
func GetBackends(c *fiber.Ctx) error {
	return c.JSON(backendList)
}

// Optional: remove by index
func RemoveBackend(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 || id >= len(backendList) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid index"})
	}
	backendList = append(backendList[:id], backendList[id+1:]...)
	return c.JSON(fiber.Map{"message": "Backend removed"})
}
