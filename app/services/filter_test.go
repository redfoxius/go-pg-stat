package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat/models"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestBuildFilter(t *testing.T) {
	app := fiber.New()

	app.Get("/api/stat/get", func(c *fiber.Ctx) error {
		testFilter := BuildFilter(c)

		sampleFilter := models.StatFilter{
			Type:  "select",
			Sort:  "fast",
			Page:  1,
			Limit: 10,
		}

		assert.Equal(t, testFilter, sampleFilter)

		return nil
	})

	req := httptest.NewRequest(fiber.MethodGet, "/api/stat/get=fast&type=select", nil)

	_, err := app.Test(req)
	assert.Equal(t, nil, err, "app.Test(req)")
}

func TestBuildFilterFixed(t *testing.T) {
	app := fiber.New()

	app.Get("/api/stat/get", func(c *fiber.Ctx) error {
		testFilter := BuildFilter(c)

		sampleFilter := models.StatFilter{
			Type:  "",
			Sort:  "slow",
			Page:  2,
			Limit: 5,
		}

		assert.Equal(t, testFilter, sampleFilter)

		return nil
	})

	req := httptest.NewRequest(fiber.MethodGet, "/api/stat/get=qqq&type=www&page=2&limit=5", nil)

	_, err := app.Test(req)
	assert.Equal(t, nil, err, "app.Test(req)")
}
