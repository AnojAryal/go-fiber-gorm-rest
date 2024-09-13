package initializers

import "github.com/anojaryal/fiber-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{})
}
