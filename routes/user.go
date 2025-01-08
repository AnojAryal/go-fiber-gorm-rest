package routes

import (
	"errors"

	"github.com/anojaryal/fiber-api/initializers"
	"github.com/anojaryal/fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// see this as the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Password:  userModel.Password,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	// Parse the request body into the user model
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON("Failed to parse request body")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON("Failed to hash password")
	}

	// Update the user model's password field with the hashed password
	user.Password = string(hashedPassword)

	if result := initializers.DB.Create(&user); result.Error != nil {
		return c.Status(500).JSON("Failed to create user")
	}

	responseUser := CreateResponseUser(user)
	return c.Status(201).JSON(responseUser)
}

// Get all users
func GetUser(c *fiber.Ctx) error {
	users := []models.User{}

	initializers.DB.Find(&users)
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	initializers.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user doesnot exist")
	}
	return nil
}

// Get a single user by id
func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Invalid user Id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// Update a user
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Invalid user Id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	initializers.DB.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// Delete a user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Invalid user Id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted User")
}
