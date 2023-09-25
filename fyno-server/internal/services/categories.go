package services

import (
	"fyno/server/internal/models"
)

type categoryService struct {
	categoryRepository models.CategoryRepository
}

func NewCategoryService(cr models.CategoryRepository) models.CategoryService {
	return &categoryService{
		categoryRepository: cr,
	}
}

func (cs *categoryService) GetAllCategories() ([]models.Category, error) {
	return cs.categoryRepository.GetAll()
}
