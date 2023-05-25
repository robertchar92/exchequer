package postgres

import (
	"exchequer/models"
	"exchequer/utils/errors"
	"log"

	"gorm.io/gorm"
)

func (r *Repository) Update(user *models.User, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(user).Error
	if err != nil {
		log.Println("error-update-user:", err)
		return errors.ErrUnprocessableEntity
	}

	return nil
}
