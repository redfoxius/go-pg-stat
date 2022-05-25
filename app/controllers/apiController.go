package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/dto"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat"
	"github.com/redfoxius/go-pg-stat/app/services"
)

type ApiController interface {
	GetRecords(c *fiber.Ctx) error
}

type controller struct {
	statRepo stat.StatRepository
}

func InitController(repo stat.StatRepository) ApiController {
	return &controller{statRepo: repo}
}

func (k *controller) GetRecords(c *fiber.Ctx) error {
	filters := services.BuildFilter(c)

	statInfo, err := k.statRepo.GetStat(context.Background(), filters)

	if err != nil {
		return c.JSON(dto.ErrorResponse{
			Status:  `error`,
			Message: err.Error(),
		})
	}

	return c.JSON(dto.Response{
		Result: statInfo,
	})
}
