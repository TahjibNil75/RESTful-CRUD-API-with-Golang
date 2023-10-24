package main

import (
	"tahjib75/restful-crud-api/config"
	"tahjib75/restful-crud-api/controller"
	"tahjib75/restful-crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDB()

	repository := controller.Repository{DB: db}
	controller := controller.Controller{Repository: repository}
	router := gin.Default()

	r := routes.Routes{Engine: router}
	r.Admin(controller)
	router.Run(":8080")
}
