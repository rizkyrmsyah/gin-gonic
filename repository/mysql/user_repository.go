package mysql

import (
	"bytes"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepositoryI {
	return &UserRepository{conn: db}
}

func (r *UserRepository) GetAll() ([]model.UserResponse, error) {
	var query bytes.Buffer
	var result []model.UserResponse
	var queryParams []interface{}
	var err error

	query.WriteString(`SELECT * FROM users`)
	err = r.conn.Select(&result, query.String(), queryParams...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) Store(params *model.StoreUserRequest) error {
	var query bytes.Buffer
	var err error

	query.WriteString("INSERT INTO users")
	query.WriteString(`(name, email, password, phone, status, created_at, updated_at)`)
	query.WriteString(`VALUES(:name, :email, :password, :phone, :status, :created_at, :updated_at)`)

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return hashErr
	}

	queryParams := map[string]interface{}{
		"name":       params.Name,
		"email":      params.Email,
		"password":   hashedPassword,
		"phone":      params.Phone,
		"status":     "active",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	_, err = r.conn.NamedExec(query.String(), queryParams)
	if err != nil {
		return err
	}

	return nil
}
