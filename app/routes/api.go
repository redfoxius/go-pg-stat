package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redfoxius/go-pg-stat/app/controllers"
	"github.com/redfoxius/go-pg-stat/app/database"
	"github.com/redfoxius/go-pg-stat/app/middlewares"
	"github.com/redfoxius/go-pg-stat/app/repositories/stat"
)

func Setup(app *fiber.App) {
	app.Use(cors.New())
	app.Use("/api", middlewares.Auth)

	db, err := database.GetDB()

	if err != nil {
		panic("DB connection Error")
	}

	repo := stat.New(db)

	cntrl := controllers.InitController(repo)

	app.Get("/api/stat/get", cntrl.GetRecords)
}
