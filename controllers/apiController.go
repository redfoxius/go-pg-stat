package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/dto"
)

func GetRecords(c *fiber.Ctx) error {

	return c.JSON(dto.Response{})
}