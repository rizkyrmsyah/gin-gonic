package usecase

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/repository"
)

type userUsecase struct {
	userInterface repository.UserRepositoryI
}

func NewUserUseCase(userInterface repository.UserRepositoryI) UserUseCaseI {
	return &userUsecase{
		userInterface,
	}
}

func (userUC *userUsecase) GetAll() ([]model.User, error) {
	var list []model.User

	list, err := userUC.GetAll()
	if err != nil {
		return list, err
	}

	return list, nil
}
