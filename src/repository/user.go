package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/k197781/go-jwt/src/entity"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserById(id int64) (*entity.User, error)
}

func (repo *userRepository) GetUserById(id int64) (*entity.User, error) {
	user := &entity.User{}
	result := repo.db.Where("id = ?", id).Find(user)
	log.Println(user)
	if result.RecordNotFound() {
		err := errors.New("User not found.")
		return nil, err
	}
	return user, nil
}

func NewUserRepository(
	db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
