package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redfoxius/go-pg-stat/database"
	"github.com/redfoxius/go-pg-stat/routes"
	"log"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	conn := database.Connect()
	defer conn.Close(context.Background())

	database.CheckRequirements(conn)

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":13013")

	//var query string
	//var calls, meanTime int
	//
	//err = conn.QueryRow(context.Background(), "SELECT `query`, `calls`, `mean_time` FROM pg_stat_statements ORDER BY mean_time DESC;").Scan(&query, &calls, &meanTime)
	//
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(query, calls, meanTime)
}
