package main

import (
	"colibri/pkg/db"
	communities_handlers "colibri/pkg/handlers/communities"
	public_handlers "colibri/pkg/handlers/public"
	topic_handlers "colibri/pkg/handlers/topics"
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
	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	// get user
	protectedRoutes.GET("/user", user_handlers.GetUser)
	// update user
	protectedRoutes.PUT("/user", user_handlers.UpdateUser)
	// delete user
	protectedRoutes.DELETE("/user", user_handlers.DeleteUser)

	// TODO: Entities
	// COMMUNITY
	// get all
	protectedRoutes.GET("/community", communities_handlers.GetCommunities)
	// create
	protectedRoutes.POST("/community", communities_handlers.CreateCommunity)
	// update
	protectedRoutes.PUT("/community/:communityId", communities_handlers.UpdateCommunity)
	// archive
	protectedRoutes.DELETE("/community/:communityId", communities_handlers.SoftDeleteCommunity)

	// TOPIC
	// get all
	// route: GET /community/:id/topic
	protectedRoutes.GET("/community/:communityId/topic", topic_handlers.GetTopics)

	// get one (with comments)
	// route: GET /community/:community-id/topic/:topic-id
	protectedRoutes.GET("/community/:communityId/topic/:topicId", topic_handlers.GetTopic)

	// create
	// route: POST /community/:community-id/topic
	protectedRoutes.POST("/community/:communityId/topic", topic_handlers.CreateTopic)

	// update
	// route: PUT /community/:community-id/topic/:topicId

	// archive
	// route: DELETE /community/:community-id/topic/:topicId

	// COMMENT
	// create

	r.Run() // listen and serve on 0.0.0.0:8080
}
