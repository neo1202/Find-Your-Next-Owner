package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostImages struct {
	Url  string `json:"url"`
	Rank int    `json:"rank"`
}
type Post struct {
	Username   string       `json:"username"`
	AvatarURL  string       `json:"avatar_url"`
	ID         uuid.UUID    `json:"id"`
	UserID     uuid.UUID    `json:"userID"`
	Kind       string       `json:"kind"`
	Name       string       `json:"name"`
	Age        string       `json:"age"`
	Gender     string       `json:"gender"`
	Content    string       `json:"content"`
	Category   Category     `json:"category"`
	Location   Location     `json:"location"`
	PostImages []PostImages `json:"post_images"`
	CreatedAt  string       `json:"created_at"`
}

type PostHandlers interface {
	GetAllPosts(c *gin.Context)
	GetPost(c *gin.Context)
	CreatePost(c *gin.Context)
}

type PostService interface {
	GetAllPosts(uuid.UUID) ([]*Post, error)
	GetPost(uuid.UUID) (*Post, error)
	CreatePost(p *Post) (uuid.UUID, error)
	CreatePostImage(p []PostImages, postID uuid.UUID) error
	DeleteAllPosts() error
}

type PostRepository interface {
	GetAll(uuid.UUID) ([]*Post, error)
	Get(uuid.UUID) (*Post, error)
	Create(p *Post) (uuid.UUID, error)
	CreatePostImage(id uuid.UUID, p PostImages, postID uuid.UUID) error
	DeleteAll() error
}
