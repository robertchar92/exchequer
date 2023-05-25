package request

type SignUpRequest struct {
	Username string `form:"username" json:"username" binding:"required,unique=users.username,min=6,max=20"`
	Email    string `form:"email" json:"email" binding:"required,unique=users.email,email"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

type SignInRequest struct {
	UsernameOrEmail string `form:"username_or_email" json:"username_or_email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
}
