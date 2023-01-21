package main

import (
	"log"
	"nokogiri-api/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.InitRouter(app)

	log.Fatal(app.Listen(":8080"))
}
