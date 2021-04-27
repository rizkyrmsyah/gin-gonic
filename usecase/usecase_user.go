package usecase

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
)

type UserUsecase struct {
	userInterface UserUseCaseInterface
}

func NewUserUseCase(userInterface UserUseCaseInterface) UserUseCaseInterface {
	return &UserUsecase{
		userInterface,
	}
}

func (userUC *UserUsecase) GetAll() ([]model.User, error) {
	var list []model.User

	list, err := userUC.GetAll()
	if err != nil {
		return list, err
	}

	return list, nil
}
