package main

import (
	"log"

	"github.com/anojaryal/fiber-api/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectToDB()

}

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome")
}

func main() {
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
