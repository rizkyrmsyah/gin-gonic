package usecase

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/repository"
)

type User struct {
	userI repository.UserRepositoryI
}

func NewUser(userI repository.UserRepositoryI) UserUseCaseI {
	return &User{
		userI,
	}
}

func (userUC *User) GetAll() ([]model.User, error) {
	var list []model.User

	list, err := userUC.userI.GetAll()
	if err != nil {
		return list, err
	}

	return list, nil
}
