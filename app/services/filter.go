package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat/models"
	"strconv"
)

const (
	pageN  = `page`
	limitN = `limit`
	sortN  = `sort`
	typeN  = `type`

	pageD  = `1`
	limitD = `10`
	sortD  = `slow`
)

func BuildFilter(c *fiber.Ctx) models.StatFilter {
	page, err := strconv.Atoi(c.Query(pageN, pageD))

	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query(limitN, limitD))

	if err != nil {
		limit = 10
	}

	sortBy := c.Query(sortN, sortD)

	if sortBy != `fast` && sortBy != `slow` {
		sortBy = `slow`
	}

	typeOnly := c.Query(typeN, ``)

	if typeOnly != `insert` && typeOnly != `update` && typeOnly != `delete` && typeOnly != `select` {
		typeOnly = ``
	}

	return models.StatFilter{
		Type:  typeOnly,
		Sort:  sortBy,
		Page:  page,
		Limit: limit,
	}
}
