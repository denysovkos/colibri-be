package main

import (
	"colibri/pkg/db"
	public_handlers "colibri/pkg/handlers/public"
	user_handlers "colibri/pkg/handlers/users"
	"colibri/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ROUTES
	// PUBLIC
	// signup
	r.POST("/signup", public_handlers.SignUp)
	// login
	r.POST("/login", public_handlers.Login)

	// PRIVATE
	userRoutes := r.Group("/api")
	userRoutes.Use(middlewares.JwtAuthMiddleware())
	// get user
	userRoutes.GET("/user", user_handlers.GetUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
