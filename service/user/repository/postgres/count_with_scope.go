package postgres

import (
	"log"

	"exchequer/models"
	"exchequer/utils/errors"
	"exchequer/utils/request_util"
)

func (r *Repository) CountWithScope(config request_util.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.User{}).
		Scopes(config.Scopes()...).
		Count(&count).Error
	if err != nil {
		log.Println("error-count-user:", err)
		return 0, errors.CustomWrap(err)
	}

	return count, nil
}
