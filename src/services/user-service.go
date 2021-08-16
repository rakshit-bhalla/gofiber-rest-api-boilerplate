package services

import (
	"rakshit.dev/gofiber-rest-api-boilerplate/src/models"
	"rakshit.dev/gofiber-rest-api-boilerplate/src/repositories"
)

type User = models.User
type UserReq = models.UserReq

// UserService ...
type UserService interface {
	GetUser(userID string) (*User, error)
	GetUsers() ([]User, error)
	DeleteUser(userID string) (*User, error)
	CreateUser(userReq UserReq) (*User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

// NewUserService ...
func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) GetUser(userID string) (*User, error) {
	return u.userRepository.GetUser(userID)
}

func (u *userService) GetUsers() ([]User, error) {
	return u.userRepository.GetUsers()
}

func (u *userService) DeleteUser(userID string) (*User, error) {
	return u.userRepository.DeleteUser(userID)
}

func (u *userService) CreateUser(userReq UserReq) (*User, error) {
	return u.userRepository.CreateUser(userReq)
}
