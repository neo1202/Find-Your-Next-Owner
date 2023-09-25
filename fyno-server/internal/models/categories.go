package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type CategoryHandlers interface {
	GetAllCategories(c *gin.Context)
}

type CategoryService interface {
	GetAllCategories() ([]Category, error)
}

type CategoryRepository interface {
	GetAll() ([]Category, error)
}
