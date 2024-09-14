package routes

import (
	"github.com/anojaryal/fiber-api/initializers"
	"github.com/anojaryal/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	//this is not the model User see this as the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error)
	}

	initializers.DB.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)
}
