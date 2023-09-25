package services

import (
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type postService struct {
	postRepository models.PostRepository
}

func NewPostService(pr models.PostRepository) models.PostService {
	return &postService{
		postRepository: pr,
	}
}

func (ps *postService) GetAllPosts(userID uuid.UUID) ([]*models.Post, error) {
	return ps.postRepository.GetAll(userID)
}

func (ps *postService) GetPost(id uuid.UUID) (*models.Post, error) {
	return ps.postRepository.Get(id)
}

func (ps *postService) CreatePost(p *models.Post) (uuid.UUID, error) {
	return ps.postRepository.Create(p)
}

func (ps *postService) CreatePostImage(p []models.PostImages, postID uuid.UUID) error {
	for _, v := range p {
		id := uuid.New()
		err := ps.postRepository.CreatePostImage(id, v, postID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ps *postService) DeleteAllPosts() error {
	return ps.postRepository.DeleteAll()
}
