package models

import "time"

type User struct {
	ID          string     `json:"id" groups:"user,admin"`
	Username    string     `json:"username" groups:"user,admin"`
	Name        string     `json:"name" groups:"user,admin"`
	Email       string     `json:"email" groups:"user,admin"`
	LastLoginAt *time.Time `json:"last_login_at" groups:"admin"`
	CreatedAt   time.Time  `json:"created_at" groups:"admin"`
	UpdatedAt   time.Time  `json:"-"`
}
