package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/k197781/go-jwt/src/entity"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserById(id int64) (*entity.User, error)
	CreateUser(email string, name string, password string) (*entity.User, error)
}

func (repo *userRepository) GetUserById(id int64) (*entity.User, error) {
	user := &entity.User{}
	result := repo.db.Where("id = ?", id).Find(user)
	if result.RecordNotFound() {
		err := errors.New("User not found.")
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) CreateUser(email string, name string, password string) (*entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Email:    email,
		Name:     name,
		Password: string(hash),
	}
	repo.db.Create(&user)
	return user, nil
}

func NewUserRepository(
	db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
