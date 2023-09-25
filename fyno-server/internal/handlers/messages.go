package handlers

import (
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type messageHandlers struct {
	messageService   models.MessageService
	websocketService models.WebSocketService
}

func NewMessageHandlers(ms models.MessageService, ws models.WebSocketService) models.MessageHandlers {
	return &messageHandlers{
		messageService:   ms,
		websocketService: ws,
	}
}

func (mh *messageHandlers) GetAllUserGroups(c *gin.Context) {
	userID := utils.StringToUUID(c.MustGet("userID").(string))
	messages, err := mh.messageService.GetAllUserGroups(userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

type CreeateUserGroupRequest struct {
	MessagePartnerID string `json:"message_partner_id"`
}

func (mh *messageHandlers) CreateUserGroup(c *gin.Context) {
	userID := utils.StringToUUID(c.MustGet("userID").(string))
	var input *CreeateUserGroupRequest
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user group already exists
	exists, err := mh.messageService.IsUserGroupExists(userID, utils.StringToUUID(input.MessagePartnerID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.Status(http.StatusOK)
		return
	}

	err = mh.messageService.CreateUserGroup(userID, utils.StringToUUID(input.MessagePartnerID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the receiver user group already exists
	exists, err = mh.messageService.IsUserGroupExists(utils.StringToUUID(input.MessagePartnerID), userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.Status(http.StatusOK)
		return
	}

	err = mh.messageService.CreateUserGroup(utils.StringToUUID(input.MessagePartnerID), userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
