package handlers

import (
	"fyno/server/internal/models"
	"fyno/server/internal/services"
)

type Handlers struct {
	Users      models.UserHandlers
	Posts      models.PostHandlers
	Messages   models.MessageHandlers
	WebSockets models.WebSocketHandlers
	Locations  models.LocationHandlers
	Categories models.CategoryHandlers
	S3         models.S3Handlers
}

func NewHandlers(serv *services.Services) *Handlers {
	return &Handlers{
		Users:      NewUserHandlers(serv.Users),
		Posts:      NewPostHandlers(serv.Posts),
		Messages:   NewMessageHandlers(serv.Messages, serv.WebSocket),
		WebSockets: NewWebSocketHandlers(serv.Messages, serv.WebSocket),
		Locations:  NewLocationHandlers(serv.Locations),
		Categories: NewCategoryHandlers(serv.Categories),
		S3:         NewS3Handlers(serv.S3),
	}
}
