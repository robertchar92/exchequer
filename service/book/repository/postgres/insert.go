package postgres

import (
	"exchequer/models"
	"exchequer/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Insert(book *models.Book, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(book).Error
	if err != nil {
		log.Println("error-insert-book:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
