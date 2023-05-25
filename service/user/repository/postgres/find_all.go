package postgres

import (
	"exchequer/models"
	"exchequer/utils/errors"
	"exchequer/utils/request_util"
	"log"
)

func (r *Repository) FindAll(config request_util.PaginationConfig) ([]models.User, error) {

	results := make([]models.User, 0)

	err := r.db.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-user:", err)
		return nil, errors.CustomWrap(err)
	}

	return results, nil
}
