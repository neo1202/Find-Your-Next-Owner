package repositories

import (
	"database/sql"
	"fyno/server/internal/models"
)

type Repositories struct {
	Users      models.UserRepository
	Posts      models.PostRepository
	Messages   models.MessageRepository
	Locations  models.LocationRepository
	Categories models.CategoryRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:      NewUserRepository(db),
		Posts:      NewPostRepository(db),
		Messages:   NewMessageRepository(db),
		Locations:  NewLocationRepository(db),
		Categories: NewCategoryRepository(db),
	}
}
