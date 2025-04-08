// handlers.go
package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siddhu949/leanbalancer/internal/firewall"
)

// HealthCheckHandler - Exposes health check functionality
func HealthCheckHandler(ctx *fiber.Ctx) error {
	response := map[string]string{"status": "Healthy"}
	return ctx.JSON(response)
}

// GetFirewallRules - Exposes firewall rules functionality
func GetFirewallRules(ctx *fiber.Ctx) error {
	blockedIPs := firewall.GetBlockedIPs()
	return ctx.JSON(blockedIPs)
}

// BlockIP - Exposes block IP functionality
func BlockIP(ctx *fiber.Ctx) error {
	type BlockRequest struct {
		IP string `json:"ip"`
	}
	var req BlockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	firewall.BlockIP(req.IP)
	response := map[string]string{"message": "IP Blocked", "ip": req.IP}
	return ctx.JSON(response)
}
