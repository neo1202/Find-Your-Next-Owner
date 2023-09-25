package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketHandlers interface {
	WsConnection(c *gin.Context)
	IsUserConnected(c *gin.Context)
}

type WebSocketService interface {
	GetWebSocketID(c *gin.Context) string
	Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)
	AddConnection(userID uuid.UUID, id string, conn *websocket.Conn)
	RemoveConnection(userID uuid.UUID)
	GetUserConnection(userID uuid.UUID) (*websocket.Conn, error)

	ListenForMessages(conn *websocket.Conn, ch chan Message)
	HandleIncomingMessages(ch chan Message)

	IsUserConnected(userID uuid.UUID) bool
}
