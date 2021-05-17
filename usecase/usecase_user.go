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
			Id:     v.Id,
			Name:   v.Name.String,
			Phone:  v.Phone.String,
			Email:  v.Email.String,
			Dob:    v.Dob.Time.String(),
			Status: v.Status.String,
			Image:  v.Image.String,
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

func (user *User) Show(id int64) (*model.UserDetail, error) {
	resp, err := user.repo.Show(id)
	if err != nil {
		return nil, err
	}

	result := &model.UserDetail{
		Id:         resp.Id,
		Name:       resp.Name.String,
		Phone:      resp.Phone.String,
		Email:      resp.Email.String,
		Dob:        resp.Dob.Time.String(),
		Status:     resp.Status.String,
		Image:      resp.Image.String,
		LastSignIn: resp.VerifiedAt.Time,
	}

	return result, nil
}
