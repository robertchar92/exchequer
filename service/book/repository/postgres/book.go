package postgres

import (
	"gorm.io/gorm"

	"exchequer/service/book"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.Repository {
	return &Repository{
		db: db,
	}
}
