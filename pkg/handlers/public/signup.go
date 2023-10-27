package public_handlers

import (
	"colibri/pkg/auth"
	"colibri/pkg/db"
	"colibri/pkg/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpRequestBody struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	NameHandler     string `json:"nameHandler"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"backgroundImage"`
}

func SignUp(c *gin.Context) {
	var requestBody SignUpRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := db.GetDBInstance()
	user := models.User{
		Email:           requestBody.Email,
		FirstName:       requestBody.FirstName,
		LastName:        requestBody.LastName,
		NameHandler:     requestBody.NameHandler,
		Password:        auth.GenerateHashedPassword(requestBody.Password),
		Avatar:          requestBody.Avatar,
		BackgroundImage: requestBody.BackgroundImage,
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok": true,
	})
}
