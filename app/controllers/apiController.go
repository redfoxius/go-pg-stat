package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/database"
	"github.com/redfoxius/go-pg-stat/app/dto"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat"
	"github.com/redfoxius/go-pg-stat/app/services"
)

func GetRecords(c *fiber.Ctx) error {
	filters := services.BuildFilter(c)

	db, err := database.GetDB()

	if err != nil {
		return err
	}

	repo := stat.New(db)

	statInfo, err := repo.GetStat(context.Background(), filters)

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
