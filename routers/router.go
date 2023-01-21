package routers

import (
	"fmt"
	"nokogiri-api/db"

	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!!!!!!!!")
	})

	api.Get("/test", func(c *fiber.Ctx) error {
		temp := db.User{Id: "1"}
		result2 := db.Db.First(&temp)
		fmt.Println(result2)
		// test := db.Db.Take(temp)
		// fmt.Pr)
		// return c.SendString("Hello,")
		if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
			log.Fatal(result2.Error)
		}
		return c.JSON(&temp)
	})
}
