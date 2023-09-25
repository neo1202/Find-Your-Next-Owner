package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PresignedURLRequest struct {
	Key string `json:"key"`
}

type S3Handlers interface {
	CreatePresignedURL(c *gin.Context)
}

type S3Service interface {
	CreatePresignedURL(key string, userID uuid.UUID) (string, error)
}
