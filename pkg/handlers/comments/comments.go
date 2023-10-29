package comments_handlers

import (
	"colibri/pkg/auth"
	"colibri/pkg/db"
	"colibri/pkg/db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageRequestBody struct {
	Message string `json:"message"`
}

func CreateComment(c *gin.Context) {
	userId, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	topicId := c.Param("topicId")
	topicIdUint, err := strconv.ParseUint(topicId, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestBody MessageRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := db.GetDBInstance()

	message := models.Comments{
		Message: requestBody.Message,
		TopicID: uint(topicIdUint),
		UserID:  userId,
	}

	result := db.Create(&message)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": result.RowsAffected > 0,
	})
}
