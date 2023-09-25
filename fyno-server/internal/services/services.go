package services

import (
	"fyno/server/internal/models"
	"fyno/server/internal/repositories"
)

type Services struct {
	Users      models.UserService
	Posts      models.PostService
	Messages   models.MessageService
	WebSocket  models.WebSocketService
	Locations  models.LocationService
	Categories models.CategoryService
	S3         models.S3Service
}

func NewServices(repositories *repositories.Repositories) *Services {
	return &Services{
		Users:      NewUserService(repositories.Users, repositories.Posts),
		Posts:      NewPostService(repositories.Posts),
		Messages:   NewMessageService(repositories.Messages),
		WebSocket:  NewWebSocketService(NewMessageService(repositories.Messages)),
		Locations:  NewLocationService(repositories.Locations),
		Categories: NewCategoryService(repositories.Categories),

		S3: NewS3Service(),
	}
}
