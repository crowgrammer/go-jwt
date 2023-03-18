package initializers

import models "github.com/crowgrammer/go-jwt/Models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
