package services

import (
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type userService struct {
	userRepository models.UserRepository
	postRepository models.PostRepository
}

func NewUserService(ur models.UserRepository, pr models.PostRepository) models.UserService {
	return &userService{
		userRepository: ur,
		postRepository: pr,
	}
}

func (us *userService) GetUser(id uuid.UUID) (*models.User, error) {
	return us.userRepository.Get(id)
}

func (us *userService) GetUserByName(name string) (*models.User, error) {
	return us.userRepository.GetByName(name)
}

func (us *userService) CreateUser(u *models.User) (*models.User, error) {
	return us.userRepository.Create(u)
}

func (us *userService) UpdateUser(userID uuid.UUID, u *models.User) (*models.User, error) {
	return us.userRepository.Update(userID, u)
}

func (us *userService) DeleteAllUsers() error {
	return us.userRepository.DeleteAll()
}

func (us *userService) GetUserPosts(userID uuid.UUID) ([]*models.Post, error) {
	return us.postRepository.GetAll(userID)
}
