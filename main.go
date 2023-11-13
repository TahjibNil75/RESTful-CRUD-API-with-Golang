package main

import (
	"tahjib75/restful-crud-api/config"
	router "tahjib75/restful-crud-api/routes"
	"tahjib75/restful-crud-api/sessions"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	db := config.ConnectToDB()

// 	repository := controller.Repository{DB: db}
// 	controller := controller.Controller{Repository: repository}
// 	router := gin.Default()

// 	r := routes.Routes{Engine: router}
// 	r.Admin(controller)
// 	router.Run(":8080")
// }

func main() {
	// Connect to the database
	db := config.ConnectToDB()

	// Migrate the models
	config.Migrate(db)

	// Create a new Gin router
	r := gin.Default()

	// Initialize the sessions controller with the repository
	sessionsController := sessions.Controller{
		Repository: sessions.Repository{
			DB: db,
		},
	}

	// Initialize routes
	routes := router.Routes{
		Engine: r,
	}

	//  Define the routes
	routes.User(sessionsController)

	// Start the server
	r.Run(":8080")
}
