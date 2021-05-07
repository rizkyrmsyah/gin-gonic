package usecase

import (
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/repository"
)

type User struct {
	repo repository.UserRepositoryI
}

func NewUser(repo repository.UserRepositoryI) UserUseCaseI {
	return &User{
		repo,
	}
}

func (user *User) GetAll() ([]*model.User, error) {
	resp, err := user.repo.GetAll()
	if err != nil {
		panic(err)
	}

	data := make([]*model.User, 0)

	for _, v := range resp {
		result := &model.User{
			Id:    v.Id,
			Name:  v.Name.String,
			Phone: v.Phone.String,
		}
		data = append(data, result)
	}

	return data, nil
}

func (user *User) Store(params *model.StoreUserRequest) error {
	if err := user.repo.Store(params); err != nil {
		return err
	}
	return nil
}
