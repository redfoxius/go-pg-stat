package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	app := fiber.New()

	app.Use("/api", Auth)
	app.Get("/api/stat/get", func(c *fiber.Ctx) error {
		return c.SendString("some stat here")
	})

	req := httptest.NewRequest(fiber.MethodGet, "/api/stat/get", nil)
	req.Header.Set("Authorization", apiKey)

	resp, err := app.Test(req)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, nil, err, "app.Test(req)")
	assert.Equal(t, 200, resp.StatusCode, "Status code")
	assert.Equal(t, "some stat here", string(body))
}

func TestAuthFail(t *testing.T) {
	app := fiber.New()

	app.Use("/api", Auth)
	app.Get("/api/stat/get", func(c *fiber.Ctx) error {
		return c.SendString("some stat here")
	})

	req := httptest.NewRequest(fiber.MethodGet, "/api/stat/get", nil)

	resp, err := app.Test(req)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, nil, err, "app.Test(req)")
	assert.Equal(t, 401, resp.StatusCode, "Status code")
	assert.Equal(t, `{"status":"error","message":"API key is mandatory."}`, string(body))

	req.Header.Set("Authorization", "qwerty")

	resp, err = app.Test(req)

	body, err = ioutil.ReadAll(resp.Body)
	assert.Equal(t, nil, err, "app.Test(req)")
	assert.Equal(t, 401, resp.StatusCode, "Status code")
	assert.Equal(t, `{"status":"error","message":"API key is wrong."}`, string(body))
}
