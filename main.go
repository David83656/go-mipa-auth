package main

import (
	"github.com/David83656/go-mipa-auth/controllers"
	"github.com/David83656/go-mipa-auth/initializers"
	"github.com/David83656/go-mipa-auth/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	//r.POST("/cobros",controllers.PeticiónMp())
	// r.POST("/cp",controllers.Preciopostal())
	r.Run()
}
