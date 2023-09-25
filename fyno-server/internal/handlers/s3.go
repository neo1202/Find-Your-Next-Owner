package handlers

import (
	"fmt"
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type S3Handlers struct {
	s3Service models.S3Service
}

func NewS3Handlers(ss models.S3Service) models.S3Handlers {
	return &S3Handlers{
		s3Service: ss,
	}
}

func (s3h *S3Handlers) CreatePresignedURL(c *gin.Context) {
	fmt.Println("get presigned url")
	var input *models.PresignedURLRequest
	userID := c.MustGet("userID").(string)
	fmt.Println(userID)
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := s3h.s3Service.CreatePresignedURL(input.Key, utils.StringToUUID(userID))
	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
