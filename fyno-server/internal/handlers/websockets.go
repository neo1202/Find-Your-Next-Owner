package handlers

import (
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type webSocketHandlers struct {
	messageService   models.MessageService
	websocketService models.WebSocketService
}

func NewWebSocketHandlers(ms models.MessageService, ws models.WebSocketService) models.WebSocketHandlers {
	return &webSocketHandlers{
		messageService:   ms,
		websocketService: ws,
	}
}

func (wh *webSocketHandlers) WsConnection(c *gin.Context) {
	userID := utils.StringToUUID(c.Param("user_id"))

	// Try to get the WebSocket ID from the cookie
	id := wh.websocketService.GetWebSocketID(c)

	conn, err := wh.websocketService.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Add the WebSocket connection to a global map using the WebSocket ID as the key
	wh.websocketService.AddConnection(userID, id, conn)

	// Send history messages
	messages, err := wh.messageService.GetAllMessages(userID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, msg := range messages {
		if err = wh.messageService.WriteMessage(conn, msg); err != nil {
			log.Println(err)
			continue
		}

	}

	ch := make(chan models.Message)
	// Start listening for incoming messages
	go wh.websocketService.ListenForMessages(conn, ch)
	// Start handling incoming messages
	go wh.websocketService.HandleIncomingMessages(ch)
}

func (wh *webSocketHandlers) IsUserConnected(c *gin.Context) {
	userID := utils.StringToUUID(c.Param("user_id"))
	connected := wh.websocketService.IsUserConnected(userID)

	c.JSON(200, gin.H{
		"connected": connected,
	})
}
