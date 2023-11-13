package sessions

import (
	"tahjib75/restful-crud-api/models"

	"github.com/go-playground/validator/v10"
)

type AdminModelValidator struct {
	Admin struct {
		// UserName        string `json:"username" binding:"required,min=4,max=255" validator:"required,min=4,max=255"`
		// Email           string `json:"email" binding:"required,email" validator:"required,email" unique:"true"`
		UserName        string `json:"username" binding:"required,min=4,max=255" validator:"required,min=4,max=255"`
		Email           string `json:"email" binding:"required,email" validator:"required,email" unique:"true"`
		Password        string `json:"password" binding:"required,min=4,max=255" validator:"required,min=4,max=255"`
		PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	} `json:"admin" validate:"required"`
	adminModel models.Admin `json:"-"`
}

func ValidateAdminModel(adminModelValidator AdminModelValidator) error {
	validate := validator.New()
	return validate.Struct(adminModelValidator)

}

type LogInValidator struct {
	Admin struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=255"`
	} `json:"admin" validate:"required"`
	adminModel models.Admin `json:"-"`
}

func ValidateAdminLogInModel(logInValidator LogInValidator) error {
	validate := validator.New()
	return validate.Struct(logInValidator)
}
