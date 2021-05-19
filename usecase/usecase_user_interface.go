package usecase

import "github.com/rizkyrmsyah/gin-gonic/model"

type UserUseCaseI interface {
	GetAll() ([]*model.User, error)
	Store(params *model.StoreUserRequest) error
	Show(id int64) (*model.UserDetail, error)
	Delete(id int64) error
	Update(params *model.StoreUserRequest, id int64) error
}
