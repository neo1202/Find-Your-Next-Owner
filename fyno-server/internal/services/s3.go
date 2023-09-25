package services

import (
	"fmt"
	"fyno/server/internal/models"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type s3Service struct {
	s3Client *s3.S3
}

func NewS3Service() models.S3Service {
	// Create a single AWS session (we can re use this if we're uploading many files)
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Region:      aws.String(region),
	})

	if err != nil {
		log.Println("Failed to create session", err)
		return nil
	}
	log.Println("Successfully created S3 session")
	// Create S3 service client
	svc := s3.New(sess)
	return &s3Service{
		s3Client: svc,
	}
}

func (ss *s3Service) CreatePresignedURL(key string, userID uuid.UUID) (string, error) {
	fmt.Println("create presigned url", key, userID)
	req, _ := ss.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME_POSTS_IMAGES")),
		Key:    aws.String(fmt.Sprintf("%s/%s", userID, key)),
	})
	str, err := req.Presign(15 * time.Minute)

	return str, err
}
