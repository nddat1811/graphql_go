package models

import (
	"errors"
	"v1_gorm_direct/controller/repository"
	"v1_gorm_direct/graph/model"
	"v1_gorm_direct/model/convert"
	"v1_gorm_direct/model/tablemodel"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserModel(db *gorm.DB) repository.UserRepository {
	return &UserModel{
		db: db,
	}
}

type UserModel struct {
	db *gorm.DB
}

func (m *UserModel) GetAllUsers() ([]*model.User, error) {
	var users []*tablemodel.User
	m.db.Select(`*`).Find(&users)

	return convert.Users(users), nil
}

func (m *UserModel) FindUserByEmailAndPassword(email string, password string) (*model.User, error) {
	var user tablemodel.User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = m.CheckPasswordHash(password, user.Password)
	if err != nil {
		return nil, err
	}

	//check if user is deleted
	if user.DeletedAt != nil {
		return nil, errors.New("user is deleted")
	}

	return convert.User(user), nil
}

func (m *UserModel) FindUserByEmail(email string) (*model.User, error) {
	var user tablemodel.User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return convert.User(user), nil
}

func (m *UserModel) FindUserByID(id int) (*model.User, error) {
	var user tablemodel.User
	err := m.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return convert.User(user), nil
}

func (m *UserModel) CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
