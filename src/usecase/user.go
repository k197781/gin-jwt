package usecase

import (
	"github.com/k197781/go-jwt/src/entity"
	"github.com/k197781/go-jwt/src/repository"
)

type userInteractor struct {
	repository repository.UserRepository
}

type UserInteractor interface {
	GetUserById(id int64) (*entity.User, error)
}

func (i *userInteractor) GetUserById(id int64) (user *entity.User, err error) {
	user, err = i.repository.GetUserById(id)
	return
}

func NewUserInteractor(
	repository repository.UserRepository) UserInteractor {
	return &userInteractor{
		repository: repository,
	}
}
