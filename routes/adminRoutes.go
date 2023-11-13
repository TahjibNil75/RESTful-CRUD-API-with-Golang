package router

import (
	"tahjib75/restful-crud-api/sessions"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	*gin.Engine
}

func (r Routes) User(sessions sessions.Controller) {
	r.POST("/user/signup", sessions.SignupAdmin)
	r.POST("/user/signin", sessions.SignInUser)
}
