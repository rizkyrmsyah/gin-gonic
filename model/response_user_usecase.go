package model

import (
	"time"
)

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	LastSignIn time.Time `json:"last_sign_in_at"`
}
