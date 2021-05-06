package usecase

import "github.com/rizkyrmsyah/gin-gonic/model"

type UserUseCaseI interface {
	GetAll() ([]*model.User, error)
}
