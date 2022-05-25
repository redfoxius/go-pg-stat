package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/dto"
)

const apiKey = "6486f5a885b64241602b0a16ca2589de"

func Auth(c *fiber.Ctx) error {
	key := c.Get("Authorization")

	if key == "" {
		return c.Status(401).JSON(dto.ErrorResponse{
			Status:  "error",
			Message: "API key is mandatory.",
		})
	}

	if key != apiKey {
		return c.Status(401).JSON(dto.ErrorResponse{
			Status:  "error",
			Message: "API key is wrong.",
		})
	}

	return c.Next()
}
