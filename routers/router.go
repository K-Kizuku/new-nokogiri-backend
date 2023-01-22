package routers

import (
	"fmt"
	"nokogiri-api/db"
	"nokogiri-api/utils"
	"nokogiri-api/ws"

	"strings"

	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("api/v1")
	room := ws.NewRoom()

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!!!!!!!!")
	})

	api.Get("/test", func(c *fiber.Ctx) error {
		temp := db.UserInfo{Id: "1"}
		result2 := db.Db.First(&temp)
		fmt.Println(result2)
		if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
			log.Fatal(result2.Error)
		}
		return c.JSON(&temp)
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", room.ServeWs)
	// api.Post("/signUp", func(c *fiber.Ctx) error {
	// 	cruds.CreateUser()
	// })
}
func middleware(c *fiber.Ctx) {
	authorizationHeader := c.GetRespHeader("Authorization")
	if authorizationHeader != "" {
		ary := strings.Split(authorizationHeader, " ")
		if len(ary) == 2 {
			if ary[0] == "Bearer" {
				t, err := jwt.ParseWithClaims(ary[1], &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
					return utils.SigningKey, nil
				})

				if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
					userId := (*claims)["sub"].(string)
					c.Set("user_id", userId)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
	c.Next()
}
