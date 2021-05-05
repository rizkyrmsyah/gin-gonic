package mysql

import (
	"bytes"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/repository"
)

type UserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepositoryI {
	return &UserRepository{conn: db}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var query bytes.Buffer
	var result []model.User
	var queryParams []interface{}
	var err error

	query.WriteString(`SELECT * FROM users LIMIT 20`)

	err = r.conn.Select(&result, query.String(), queryParams...)

	if err != nil {
		return nil, err
	}

	return result, nil
}
