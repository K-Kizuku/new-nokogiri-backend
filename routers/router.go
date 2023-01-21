package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitRouter(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!!!!!!!!")
	})
}
