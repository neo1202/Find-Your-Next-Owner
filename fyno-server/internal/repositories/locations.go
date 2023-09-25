package repositories

import (
	"database/sql"
	"fyno/server/internal/models"
)

type locationRepository struct {
	DB *sql.DB
}

func NewLocationRepository(db *sql.DB) models.LocationRepository {
	return &locationRepository{
		DB: db,
	}
}

func (lr *locationRepository) GetAll() ([]models.Location, error) {
	query := `SELECT id, name FROM locations`
	rows, err := lr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		var l models.Location
		err := rows.Scan(&l.ID, &l.Name)
		if err != nil {
			return nil, err
		}
		locations = append(locations, l)
	}

	return locations, nil
}
