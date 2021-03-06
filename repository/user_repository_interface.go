package repository

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
)

type UserRepositoryI interface {
	GetAll() ([]model.UserResponse, error)
	Store(params *model.StoreUserRequest) error
	Show(id int64) (*model.UserResponse, error)
	Delete(id int64) error
	Update(params *model.StoreUserRequest, id int64) error
}
