package main

import (
	"log"
	"nokogiri-api/db"
	"nokogiri-api/routers"
	"nokogiri-api/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	utils.LoadEnv()
	db.InitDB()
	routers.InitRouter(app)

	log.Fatal(app.Listen(":8080"))
}
