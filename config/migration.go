package config

import (
	"tahjib75/restful-crud-api/models"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Admin{})
}
