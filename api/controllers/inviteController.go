package controllers

import (
	"log"
	"net/http"
	initializers "youtravel-api/api/initializers"
	"youtravel-api/api/models"

	"github.com/gin-gonic/gin"
)

func CreateInvite(c *gin.Context) {
	var invite models.Invite

	if err := c.ShouldBindJSON(&invite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	log.Printf("UserIDInviter: %d, UserIDInvited: %d\n", invite.UserIDInviter, invite.UserIDInvited)

	if invite.UserIDInviter == invite.UserIDInvited {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Inviter and invited user IDs cannot be the same",
		})
		return
	}

	if err := initializers.DB.Create(&invite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create invite",
		})
		return
	}

	c.JSON(http.StatusOK, invite)
}

func GetInviteByID(c *gin.Context) {
	var invite models.Invite
	inviteID := c.Param("id") // Assuming the ID is passed as a URL parameter

	// Query the database for the invite with the given ID
	if err := initializers.DB.First(&invite, inviteID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invite not found",
		})
		return
	}

	c.JSON(http.StatusOK, invite)
}

func ChangeStatusAcceptedInvited(c *gin.Context) {
	id := c.Param("id")

	var invite models.Invite
	if err := initializers.DB.First(&invite, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invite not found",
		})
		return
	}

	if invite.Status == "accepted" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invite is already accepted",
		})
		return
	}

	invite.Status = "accepted"

	if err := initializers.DB.Save(&invite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update status",
		})
		return
	}

	if err := CreateParticipation(invite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create participation record",
		})
		return
	}

	c.JSON(http.StatusOK, invite)

}

func CreateParticipation(invite models.Invite) error {
	participation := models.Participation{
		UserID:   invite.UserIDInvited,
		TravelID: invite.TravelID,
	}

	if err := initializers.DB.Create(&participation).Error; err != nil {
		return err
	}

	return nil
}

func ChangeStatusDeclinedInvited(c *gin.Context) {
	id := c.Param("id")

	var invite models.Invite
	if err := initializers.DB.First(&invite, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invite not found",
		})
		return
	}

	if invite.Status == "declined" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invite is already declined",
		})
		return
	}

	invite.Status = "declined"

	if err := initializers.DB.Save(&invite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update status",
		})
		return
	}

	c.JSON(http.StatusOK, invite)
}
