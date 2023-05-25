package request

type UserUpdateRequest struct {
	Email       *string `form:"email" json:"email" binding:"omitempty,unique=users.email,email"`
	Name        *string `form:"name" json:"name" binding:"omitempty"`
	OldPassword *string `form:"old_password" json:"old_password" binding:"omitempty,min=8"`
	NewPassword *string `form:"new_password" json:"new_password" binding:"omitempty,min=8"`
}
