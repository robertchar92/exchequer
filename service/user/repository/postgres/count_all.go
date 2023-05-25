package postgres

import (
	"log"

	"exchequer/models"
	"exchequer/utils/errors"
)

func (r *Repository) CountAll() (int64, error) {
	var count int64

	err := r.db.
		Model(&models.User{}).
		Count(&count).Error
	if err != nil {
		log.Println("error-count-all-user:", err)
		return 0, errors.CustomWrap(err)
	}

	return count, nil
}
