package model

import (
	"time"
)

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Dob        string    `json:"dob"`
	Status     string    `json:"status"`
	Image      string    `json:"image"`
	LastSignIn time.Time `json:"last_sign_in_at"`
}

type UserDetail struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Gender     bool      `json:"gender"`
	Email      string    `json:"email"`
	Dob        string    `json:"dob"`
	Status     string    `json:"status"`
	Image      string    `json:"image"`
	LastSignIn time.Time `json:"last_sign_in_at"`
	VetifiedAt time.Time `json:"verified_at"`
}
