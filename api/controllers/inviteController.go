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
