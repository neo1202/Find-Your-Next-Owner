package repositories

import (
	"database/sql"
	"fyno/server/internal/models"
)

type categoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) models.CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (cr *categoryRepository) GetAll() ([]models.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := cr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
