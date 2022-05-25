package routes

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	os.Setenv("PG_HOST", "127.0.0.1")
	os.Setenv("PG_PORT", "5432")
	os.Setenv("PG_USER", "postgres")
	os.Setenv("PG_PASSW", "postgres")
	os.Setenv("PG_DBNAME", "dvdrental")

	app := fiber.New()

	Setup(app)
}
