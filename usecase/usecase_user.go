package usecase

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/repository"
)

type UserUseCase struct {
	userInterface repository.UserRepositoryI
}

func NewUserUseCase(userInterface repository.UserRepositoryI) UserUseCaseI {
	return &UserUseCase{
		userInterface,
	}
}

func (userUC *UserUseCase) GetAll() ([]model.User, error) {
	var list []model.User

	list, err := userUC.GetAll()
	if err != nil {
		return list, err
	}

	return list, nil
}
