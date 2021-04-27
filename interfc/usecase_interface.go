package interfc

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
)

type UserInterface interface {
	FindAll() ([]model.User, error)
}
