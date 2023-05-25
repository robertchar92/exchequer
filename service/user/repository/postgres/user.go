package postgres

import (
	"gorm.io/gorm"

	"exchequer/service/user"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &Repository{
		db: db,
	}
}
