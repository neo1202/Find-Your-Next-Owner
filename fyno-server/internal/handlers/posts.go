package handlers

import (
	"fmt"
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandlers struct {
	postService models.PostService
}

func NewPostHandlers(ps models.PostService) models.PostHandlers {
	return &postHandlers{
		postService: ps,
	}
}

func (ph *postHandlers) GetAllPosts(c *gin.Context) {
	posts, err := ph.postService.GetAllPosts(uuid.Nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (ph *postHandlers) GetPost(c *gin.Context) {
	fmt.Println("GetPost")
	postID := c.Param("id")
	fmt.Println("postID", postID)
	post, err := ph.postService.GetPost(utils.StringToUUID(postID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (ph *postHandlers) CreatePost(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	var input *models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = utils.StringToUUID(userID)
	fmt.Println("input", input)
	postID, err := ph.postService.CreatePost(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ph.postService.CreatePostImage(input.PostImages, input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"postID": postID})
}
