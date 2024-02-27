package initializers

import "github.com/edwinnambaje/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}