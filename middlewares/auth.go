package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/dto"
)

const apiKey = "6486f5a885b64241602b0a16ca2589de"

func Auth(c *fiber.Ctx) error {
	apiKey := c.Get("Authorization")

	if apiKey == "" {
		return c.Status(401).JSON(dto.ErrorResponse{
			Status:  "error",
			Message: "API key is mandatory.",
		})
	}

	if apiKey != apiKey {
		return c.Status(401).JSON(dto.ErrorResponse{
			Status:  "error",
			Message: "API key is wrong.",
		})
	}

	return c.Next()
}
