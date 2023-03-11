package initializers

import "github.com/David83656/go-mipa-auth/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
