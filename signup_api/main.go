package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/initializers"
	"main.go/middleware"
)

func init() {
	initializers.LoadEnvs()
	initializers.CreateDB()
}

func main() {
	r := gin.Default()

	r.POST("/auth/signup", controllers.CreateUser)
	r.POST("/auth/login", controllers.Login)
	r.GET("/user/profile", middleware.CheckAuth, controllers.GetUserProfile)

	r.Run()
}
