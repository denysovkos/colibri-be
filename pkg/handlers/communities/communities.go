package communities_handlers

import (
	"colibri/pkg/auth"
	"colibri/pkg/db"
	"colibri/pkg/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommunityRequestBody struct {
	Name            string `json:"name"`
	BackgroundImage string `json:"backgroundImage"`
	Public          string `json:"public"`
}

func CreateCommunity(c *gin.Context) {
	// userId := c.Params.ByName("id") use this when get user by id
	userId, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var requestBody CommunityRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	community := models.Community{
		Name:            requestBody.Name,
		BackgroundImage: requestBody.BackgroundImage,
		OwnerID:         userId,
	}

	db := db.GetDBInstance()

	result := db.Create(&community)
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

func GetCommunities(c *gin.Context) {
	// userId := c.Params.ByName("id") use this when get user by id
	userId, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDBInstance()

	var communities []models.Community
	// TODO: Add pagination
	db.Debug().Where("owner_id = ?", userId).Find(&communities)

	c.JSON(http.StatusOK, gin.H{
		"communities": communities,
	})
}

func UpdateCommunity(c *gin.Context) {
	userId, err := auth.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDBInstance()

	var community models.Community
	communityID := c.Param("communityId")

	// Check if the community exists
	result := db.Where("id = ?", communityID).Where("owner_id = ?", userId).First(&community)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Community not found"})
		return
	}

	// Parse the request body into a struct representing the fields to update
	var updateCommunity models.Community
	if err := c.ShouldBindJSON(&updateCommunity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the community fields
	community.Name = updateCommunity.Name
	community.BackgroundImage = updateCommunity.BackgroundImage
	community.Public = updateCommunity.Public

	// Save the updated community
	result = db.Save(&community)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update community"})
		return
	}

	c.JSON(http.StatusOK, community)
}

func SoftDeleteCommunity(c *gin.Context) {
	var community models.Community
	communityID := c.Param("communityId")

	db := db.GetDBInstance()
	// Check if the community exists
	result := db.First(&community, communityID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Community not found"})
		return
	}

	// Soft delete the community (set DeletedAt field)
	result = db.Delete(&community)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete community"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Community deleted successfully"})
}
