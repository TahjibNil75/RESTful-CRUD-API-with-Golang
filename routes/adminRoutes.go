package routes

import (
	"tahjib75/restful-crud-api/controller"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	*gin.Engine
}

func (r Routes) Admin(controller controller.Controller) {
	r.POST("/admin/signup", controller.SignupAdmin)
}
