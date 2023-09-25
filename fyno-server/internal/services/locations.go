package services

import (
	"fyno/server/internal/models"
)

type locationService struct {
	locationRepository models.LocationRepository
}

func NewLocationService(lr models.LocationRepository) models.LocationService {
	return &locationService{
		locationRepository: lr,
	}
}

func (ls *locationService) GetAllLocations() ([]models.Location, error) {
	return ls.locationRepository.GetAll()
}
