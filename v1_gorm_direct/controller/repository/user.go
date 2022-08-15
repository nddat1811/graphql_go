package repository

import (
	"v1_gorm_direct/graph/model"
)

type UserRepository interface {
	GetAllUsers() ([]*model.User, error)
	FindUserByEmailAndPassword(email string, password string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(id int) (*model.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) error
}
