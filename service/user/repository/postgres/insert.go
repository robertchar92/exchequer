package postgres

import (
	"exchequer/models"
	"exchequer/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Insert(user *models.User, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Create(user).Error
	if err != nil {
		log.Println("error-insert-user:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
