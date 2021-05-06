package model

import (
	"database/sql"
)

type UserResponse struct {
	Id                 int64          `db:"id"`
	Name               sql.NullString `db:"name"`
	Email              sql.NullString `db:"email"`
	Password           sql.NullString `db:"password"`
	ResetPasswordToken sql.NullString `db:"reset_password_token"`
	Phone              sql.NullString `db:"phone"`
	Gender             sql.NullBool   `db:"gender"`
	Dob                sql.NullTime   `db:"dob"`
	Status             sql.NullString `db:"status"`
	Image              sql.NullString `db:"image"`
	LastSignIn         sql.NullTime   `db:"last_sign_in_at"`
	VerifiedAt         sql.NullTime   `db:"verified_at"`
	CreatedAt          sql.NullTime   `db:"created_at"`
	UpdatedAt          sql.NullTime   `db:"updated_at"`
}
