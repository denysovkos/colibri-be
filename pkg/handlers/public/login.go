package public_handlers

import (
	"colibri/pkg/auth"
	"colibri/pkg/db"
	"colibri/pkg/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserJson struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	NameHandler     string `json:"nameHandler"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"backgroundImage"`
}

func Login(c *gin.Context) {
	var requestBody LoginRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := db.GetDBInstance()
	var user models.User
	if err := db.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong credentials",
		})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	currentUser := UserJson{
		Email:           user.Email,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		NameHandler:     user.NameHandler,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  currentUser,
		"token": token,
	})
}
