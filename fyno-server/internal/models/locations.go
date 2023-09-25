package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Location struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type LocationHandlers interface {
	GetAllLocations(c *gin.Context)
}
type LocationService interface {
	GetAllLocations() ([]Location, error)
}

type LocationRepository interface {
	GetAll() ([]Location, error)
}
