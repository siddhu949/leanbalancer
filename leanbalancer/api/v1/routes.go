// routes.go
package v1

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1") // Base route

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "API v1 Running"})
	})

	v1.Get("/health", HealthCheckHandler)
	v1.Get("/firewall", GetFirewallRules)
	v1.Post("/firewall/block", BlockIP)
}
