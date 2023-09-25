package handlers

import (
	"fyno/server/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandlers struct {
	categoryService models.CategoryService
}

func NewCategoryHandlers(cs models.CategoryService) models.CategoryHandlers {
	return &categoryHandlers{
		categoryService: cs,
	}
}

func (ch *categoryHandlers) GetAllCategories(c *gin.Context) {
	categories, err := ch.categoryService.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}
