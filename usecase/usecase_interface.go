package usecase

import "github.com/rizkyrmsyah/gin-gonic/model"

type UseCaseI interface {
	FindAll() ([]model.User, error)
}
