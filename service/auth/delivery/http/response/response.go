package response

import "exchequer/models"

type AuthResponse struct {
	Token string      `json:"token" groups:"user"`
	User  models.User `json:"user" groups:"user"`
}
