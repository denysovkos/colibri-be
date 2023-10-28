package topic_handlers

import (
	"colibri/pkg/auth"
	"colibri/pkg/db"
	"colibri/pkg/db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TopicRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetTopics(c *gin.Context) {
	// TODO: Add private users topics
	_, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	communityID := c.Param("communityId")
	db := db.GetDBInstance()

	var topics []models.Topic
	db.Debug().Where("community_id = ?", communityID).Find(&topics)

	c.JSON(http.StatusOK, gin.H{
		"topics": topics,
	})
}

func GetTopic(c *gin.Context) {
	// TODO: Add private users topics
	_, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	communityId := c.Param("communityId")
	topicId := c.Param("topicId")

	db := db.GetDBInstance()

	var topic models.Topic
	db.Debug().
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Order("comments.created_at ASC")
		}).
		Preload("Comments.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, name_handler, avatar")
		}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name, name_handler, avatar")
		}).
		Where("community_id = ?", communityId).
		Where("id = ?", topicId).Find(&topic)

	c.JSON(http.StatusOK, gin.H{
		"topic": topic,
	})
}

func CreateTopic(c *gin.Context) {
	userId, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	communityId := c.Param("communityId")
	communityIdUint, err := strconv.ParseUint(communityId, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestBody TopicRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := db.GetDBInstance()

	topic := models.Topic{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		CommunityId: uint(communityIdUint),
		OwnerID:     userId,
	}

	result := db.Create(&topic)
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
