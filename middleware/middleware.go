package middleware

import (
	"net/http"
	"tahjib75/restful-crud-api/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Retrieve JWT from cookie
		tokenString, err := ctx.Cookie("AdminJwt")
		if tokenString == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Request does not contain access token",
			})
			ctx.Abort()
			return
		}
		// Use the original err variable for ValidateToken
		err = utils.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid access token",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
