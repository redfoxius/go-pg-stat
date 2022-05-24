package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redfoxius/go-pg-stat/controllers"
	"github.com/redfoxius/go-pg-stat/middlewares"
)

func Setup(app *fiber.App) {
	app.Use(cors.New())
	app.Use("/api", middlewares.Auth)

	app.Get("/api/stat/get", controllers.GetRecords)
}
