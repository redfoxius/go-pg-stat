package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redfoxius/go-pg-stat/app/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	app := fiber.New()

	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
