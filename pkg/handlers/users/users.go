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
	// maybe will be required .Preload("Communities")
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

func UpdateUser(c *gin.Context) {
	userID, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDBInstance()

	// Check if the user exists
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		// User not found, return an error
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Parse the request body into a struct representing the fields to update
	var updateUser struct {
		Email           string `json:"email"`
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		NameHandler     string `json:"nameHandler"`
		Password        string `json:"password"`
		Avatar          string `json:"avatar"`
		BackgroundImage string `json:"backgroundImage"`
	}
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.NameHandler != "" {
		user.NameHandler = updateUser.NameHandler
	}
	// TODO: Implement later change password feature
	// if updateUser.Password != "" {
	// 	user.Password = updateUser.Password
	// }
	if updateUser.Avatar != "" {
		user.Avatar = updateUser.Avatar
	}
	if updateUser.BackgroundImage != "" {
		user.BackgroundImage = updateUser.BackgroundImage
	}

	// Update the user in the database
	result = db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	currentUser := public_handlers.UserJson{
		Email:           user.Email,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		NameHandler:     user.NameHandler,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
	}

	c.JSON(http.StatusOK, gin.H{"user": currentUser})
}

func DeleteUser(c *gin.Context) {
	userID, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDBInstance()

	// Check if the user exists
	var user models.User
	result := db.First(&user, userID)
	if result.Error != nil {
		// User not found, return an error
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user from the database
	result = db.Delete(&user)
	if result.Error != nil {
		// Handle the delete error, e.g., return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
