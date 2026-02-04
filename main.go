package main

import (
	"github.com/custordev/shortener-api/controllers"
	"github.com/custordev/shortener-api/initializers"
	"github.com/custordev/shortener-api/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
    users:= v1.Group("/users")
	{
		users.POST("",controllers.SignUpWithToken)
		users.POST("/login",controllers.LoginWithToken)
		users.GET("/validate",middleware.RequireAuthWithToken,controllers.Validate)
	}

   links:=v1.Group("/links")

  {
	links.POST("/links")
  }
router.Run()

	println("Hello Go world")

}