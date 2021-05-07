package model

import "time"

type StoreUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
