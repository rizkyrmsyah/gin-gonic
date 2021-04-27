package mysql

import (
	"bytes"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyrmsyah/gin-gonic/model"
)

type UserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		conn: db,
	}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var query bytes.Buffer
	var result []model.User

	query.WriteString("`SELECT * FROM users WHERE deleted_at IS NULL LIMIT 20`")

	return result, nil
}
