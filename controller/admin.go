package controller

import (
	"net/http"
	"tahjib75/restful-crud-api/models"
	"tahjib75/restful-crud-api/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Repository Repository
}

func (c Controller) SignupAdmin(ctx *gin.Context) {
	var signUpPaylaod AdminModelValidator
	if err := ctx.ShouldBindJSON(&signUpPaylaod); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "binding json faild",
			"data":    "error ",
		})
		return
	}

	if signUpPaylaod.adminModel.Password != signUpPaylaod.adminModel.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "password do not match",
			"data":    "error ",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(signUpPaylaod.adminModel.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "failed to hash password",
			"data":    "error ",
		})
		return
	}
	now := time.Now()
	newAdmin := models.Admin{
		UserName:  signUpPaylaod.adminModel.UserName,
		Email:     signUpPaylaod.adminModel.Email,
		Password:  hashedPassword,
		IsAdmin:   true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := c.Repository.saveAdmin(&newAdmin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "failed to save admin",
			"data":    "error ",
		})
		return
	}

	adminResponse := models.AdminResponse{
		ID:        newAdmin.Uid,
		UserName:  newAdmin.UserName,
		Email:     newAdmin.Email,
		IsAdmin:   newAdmin.IsAdmin,
		CreatedAt: newAdmin.CreatedAt,
		UpdatedAt: newAdmin.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"admin":  adminResponse,
	})
}
