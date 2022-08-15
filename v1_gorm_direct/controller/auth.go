package controller

import (
	"fmt"
	"v1_gorm_direct/controller/repository"
	"v1_gorm_direct/controller/service"
	"v1_gorm_direct/graph/model"
)

type AuthControllerInterface interface {
	Login(model.UserLogin) *model.LoginResponse
}

func NewAuthController(
	userRepository repository.UserRepository,
	authService service.AuthService,
) *AuthController {
	return &AuthController{
		userRepository: userRepository,
		authService:    authService,
	}
}

type AuthController struct {
	userRepository repository.UserRepository
	authService    service.AuthService
}

func (a *AuthController) Login(userLogin model.UserLogin) *model.LoginResponse {
	user, _ := a.userRepository.FindUserByEmailAndPassword(userLogin.Email, userLogin.Password)

	token, _ := a.authService.GenerateToken(*user)

	loginResponse := &model.LoginResponse{
		Token: token,
		ID:    fmt.Sprintf(user.ID),
	}

	return loginResponse
}
