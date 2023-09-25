package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Sender    uuid.UUID `json:"sender"`
	Receiver  uuid.UUID `json:"receiver"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Type      string    `json:"type"`
}

type MessageHandlers interface {
	GetAllUserGroups(c *gin.Context)
	CreateUserGroup(c *gin.Context)
}

type MessageService interface {
	GetAllMessages(userID uuid.UUID) ([]Message, error)
	CreateMessage(m *Message) (uuid.UUID, error)
	WriteMessage(*websocket.Conn, Message) error
	GetAllUserGroups(userID uuid.UUID) ([]User, error)
	CreateUserGroup(userID uuid.UUID, messagePartnerID uuid.UUID) error
	IsUserGroupExists(userID uuid.UUID, messagePartnerID uuid.UUID) (bool, error)
}

type MessageRepository interface {
	GetAll(userID uuid.UUID) ([]Message, error)
	Create(m *Message) (uuid.UUID, error)
	GetAllUserGroups(userID uuid.UUID) ([]User, error)
	CreateUserGroup(userID uuid.UUID, messagePartnerID uuid.UUID) error
	UpdateUserGroup(userID uuid.UUID, messagePartnerID uuid.UUID) error
	IsUserGroupExists(userID uuid.UUID, messagePartnerID uuid.UUID) (bool, error)
}
