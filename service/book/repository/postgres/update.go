package postgres

import (
	"exchequer/models"
	"exchequer/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Update(book *models.Book, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(book).Error
	if err != nil {
		log.Println("error-update-book:", err)
		return errors.ErrUnprocessableEntity
	}

	return nil
}
