package main

import (
	"user-api/config"
	"user-api/routes"
	"user-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	db := config.DB

    // AutoMigrate will create the "users" table based on the User struct
    db.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.SetupUserRoutes(r)

	r.Run(":8080")
}
