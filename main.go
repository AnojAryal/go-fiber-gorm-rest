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
	//user routes
	app.Post("/api/create-user", routes.CreateUser)
	app.Get("/api/users", routes.GetUser)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Put("/api/update-user/:id", routes.UpdateUser)
	app.Delete("/api/delete-user/:id", routes.DeleteUser)
	//product routes
	app.Post("/api/products", routes.CreateProduct)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
