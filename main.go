package main

import (
	"github.com/gin-gonic/gin"
	"goku/controllers"
	"goku/models"
	"os"
)

func main() {
	router := gin.Default()

	models.OpenConnection()
	// models.DB.Create(&models.Account{Email: "bill@gates.com", Password: "pass"})

	router.GET("/accounts", controllers.GetAccounts)
	router.GET("/accounts/:id", controllers.GetAccount)
	router.POST("/accounts", controllers.PostAccount)
	router.PUT("/accounts/:id", controllers.UpdateAccount)
	router.DELETE("/accounts/:id", controllers.DeleteAccount)

	router.POST("/login", controllers.Login)

	router.Run("0.0.0.0:" + os.Getenv("DOCKER_PORT"))
}
