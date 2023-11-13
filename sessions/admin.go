package sessions

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

	token, refresh_token, err := utils.CreateAccessToken(newAdmin.Email, newAdmin.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate access token",
		})
		return
	}

	newAdmin.Token = &token
	newAdmin.RefreshToken = &refresh_token

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

func (c Controller) SignInUser(ctx *gin.Context) {
	var signInData LogInValidator
	if err := ctx.ShouldBindJSON(&signInData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Username or Password",
		})
		return
	}

	admin, err := c.Repository.FindOne(signInData.adminModel.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to Find User",
		})
		return
	}
	//TODO: correct this part
	if err := utils.VerifyPassword(admin.Password, signInData.adminModel.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong password!, Please try again!",
		})
		return
	}
	// Password verification succeeded, proceed with further actions
	// For example, generate and return a JWT token
	//token, refreshToken, err := utils.UpdateAllToken(user.Token, false)
	token, err := utils.UpdateAllTokens(*admin.Token, false)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	// Set the "AdminJwt" cookie with a 6-hour expiration time
	ctx.SetCookie("AdminJwt", token, 3600*6, "/", "", false, true)

	// Set a separate "LoggedIn" cookie to track user login status
	ctx.SetCookie("LoggedIn", "true", 3600*6, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})

}
