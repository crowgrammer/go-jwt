package main

import (
	controllers "github.com/crowgrammer/go-jwt/Controllers"
	initializers "github.com/crowgrammer/go-jwt/Initializers"
	middleware "github.com/crowgrammer/go-jwt/Middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}
func main() {

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
