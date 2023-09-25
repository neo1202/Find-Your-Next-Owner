package services

import (
	"fmt"
	"fyno/server/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type messageService struct {
	messageRepository models.MessageRepository
}

func NewMessageService(mr models.MessageRepository) models.MessageService {
	return &messageService{
		messageRepository: mr,
	}
}

func (m *messageService) WriteMessage(conn *websocket.Conn, msg models.Message) error {
	fmt.Println("msg", msg)
	return conn.WriteJSON(msg)
}

func (ms *messageService) CreateMessage(m *models.Message) (uuid.UUID, error) {
	ms.messageRepository.UpdateUserGroup(m.Sender, m.Receiver)
	ms.messageRepository.UpdateUserGroup(m.Receiver, m.Sender)
	return ms.messageRepository.Create(m)
}

func (ms *messageService) GetAllMessages(userID uuid.UUID) ([]models.Message, error) {
	return ms.messageRepository.GetAll(userID)
}

func (ms *messageService) GetAllUserGroups(userID uuid.UUID) ([]models.User, error) {
	return ms.messageRepository.GetAllUserGroups(userID)
}

func (ms *messageService) CreateUserGroup(userID uuid.UUID, messagePartnerID uuid.UUID) error {
	return ms.messageRepository.CreateUserGroup(userID, messagePartnerID)
}

func (ms *messageService) IsUserGroupExists(userID uuid.UUID, messagePartnerID uuid.UUID) (bool, error) {
	// UpdateUserGroup is used to update the user group updated_at field
	ms.messageRepository.UpdateUserGroup(userID, messagePartnerID)
	return ms.messageRepository.IsUserGroupExists(userID, messagePartnerID)
}
