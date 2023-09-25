package handlers

import (
	"fyno/server/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type locationHandlers struct {
	locationService models.LocationService
}

func NewLocationHandlers(ls models.LocationService) models.LocationHandlers {
	return &locationHandlers{
		locationService: ls,
	}
}

func (lh *locationHandlers) GetAllLocations(c *gin.Context) {
	locations, err := lh.locationService.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"locations": locations})
}
