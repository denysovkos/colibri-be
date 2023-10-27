package user_handlers

import (
	"colibri/pkg/auth"
	"colibri/pkg/db"
	"colibri/pkg/db/models"
	public_handlers "colibri/pkg/handlers/public"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// userId := c.Params.ByName("id") use this when get user by id
	userId, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDBInstance()

	var user models.User
	db.First(&user, userId)

	currentUser := public_handlers.UserJson{
		Email:           user.Email,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		NameHandler:     user.NameHandler,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
	}

	c.JSON(http.StatusOK, gin.H{
		"user": currentUser,
	})
}
