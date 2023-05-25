package user

import (
	"exchequer/models"
	"exchequer/service/user/delivery/http/request"
)

type Usecase interface {
	Show(userID string, unscoped ...bool) (*models.User, error)
	Update(userID string, request request.UserUpdateRequest) (*models.User, error)
}
