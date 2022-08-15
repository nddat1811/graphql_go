package controller

import (
	"v1_gorm_direct/controller/repository"
	"v1_gorm_direct/graph/model"
)

type UserControllerInterface interface {
	GetAllUsers() []*model.User
}

func NewUserController(
	userRepository repository.UserRepository,
) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

type UserController struct {
	userRepository repository.UserRepository
}

func (u *UserController) GetAllUsers() []*model.User {
	t, _ := u.userRepository.GetAllUsers()
	return t
}

