package main

import (
	"github.com/David83656/go-mipa-auth/controllers"
	"github.com/David83656/go-mipa-auth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	// r.POST("/signup", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Pong",
	// 	})
	// })
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.Run()
}
