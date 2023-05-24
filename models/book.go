package models

import "time"

type Book struct {
	ID        string    `json:"id" groups:"user,admin"`
	UserID    string    `json:"user_id" groups:"user,admin"`
	Username  string    `json:"username" groups:"user,admin"`
	Name      string    `json:"name" groups:"user,admin"`
	Balance   uint64    `json:"balance" groups:"user,admin"`
	CreatedAt time.Time `json:"created_at" groups:"user,admin"`
	UpdatedAt time.Time `json:"-"`
}
