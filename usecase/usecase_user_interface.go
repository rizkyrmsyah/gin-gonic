package usecase

import "github.com/rizkyrmsyah/gin-gonic/model"

type UserUseCaseInterface interface {
	GetAll() ([]model.User, error)
}
