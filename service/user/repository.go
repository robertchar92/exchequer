package user

import (
	"exchequer/models"
	"exchequer/utils/request_util"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(user *models.User, tx *gorm.DB) error
	Update(user *models.User, tx *gorm.DB) error
	FindAll(config request_util.PaginationConfig) ([]models.User, error)
	FindByID(userID string, unscoped ...bool) (*models.User, error)
	FindByUsername(username string, unscoped ...bool) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	CountWithScope(config request_util.PaginationConfig) (int64, error)
	CountAll() (int64, error)
}
