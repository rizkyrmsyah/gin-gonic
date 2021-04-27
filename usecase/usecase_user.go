package usecase

import (
	useCaseI "github.com/rizkyrmsyah/gin-gonic/interfc"
	userI "github.com/rizkyrmsyah/gin-gonic/interfc"
	"github.com/rizkyrmsyah/gin-gonic/model"
)

type UserUsecase struct {
	userInterface userI.UserInterface
}

func NewUserUseCase(userInterface useCaseI.UserInterface) useCaseI.UserUseCaseI {
	return &UserUsecase{
		userInterface,
	}
}

func (userUC *UserUsecase) FindAll() ([]model.User, error) {
	var list []model.User

	list, err := userUC.FindAll()
	if err != nil {
		return list, err
	}

	return list, nil
}
