package repository

import "github.com/rizkyrmsyah/gin-gonic/model"

type UserRepositoryI interface {
	GetAll() ([]model.User, error)
}