package interfc

import "github.com/rizkyrmsyah/gin-gonic/model"

type UserUseCaseI interface {
	FindAll() ([]model.User, error)
}
