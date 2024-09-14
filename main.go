package main

import (
	"log"

	"github.com/anojaryal/fiber-api/initializers"
	"github.com/anojaryal/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	initializers.LoadEnvVariables()
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome")
}

func setupRoutes(app *fiber.App) {
	//welcome api
	app.Get("/api", welcome)
	app.Post("/api/create-user", routes.CreateUser)

}

func main() {
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
