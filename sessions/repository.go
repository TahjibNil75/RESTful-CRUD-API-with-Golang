package sessions

import (
	"errors"
	"tahjib75/restful-crud-api/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) saveAdmin(admin *models.Admin) error {
	// Check if the admin already exists by email
	var existingAdmin models.Admin
	result := r.DB.Where("email=?", admin.Email).First(&existingAdmin)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("admin already exists")
	}

	// The admin does not exist, so we can save the new admin
	err := r.DB.Create(admin).Error
	return err
}

func (r Repository) FindOne(condition interface{}) (models.Admin, error) {
	var adminModel models.Admin
	err := r.DB.Where(condition).First(&adminModel).Error
	return adminModel, err
}
