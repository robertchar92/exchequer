package http

import (
	"github.com/gin-gonic/gin"

	"exchequer/app/middleware"
	"exchequer/service/auth"
	//"gitlab.com/depatu/core/utils/helpers"
)

type Handler struct {
	authUsecase auth.Usecase
}

func New(authUC auth.Usecase) *Handler {
	return &Handler{
		authUsecase: authUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	userRoute := r.Group("/auth")
	{
		userRoute.POST("/signup", h.SignUp)
		userRoute.POST("/signin", h.SignIn)

		// userRoute.POST("/forget-password", h.ForgetPassword)
		// userRoute.GET("/reset-password/:token", h.ValidateResetPasswordToken)
		// userRoute.POST("/reset-password", h.ResetPassword)

		// refresh jwt token
		// userRoute.GET("/refresh", m.RefreshHandle())
	}

	// adminRoute := r.Group("/admin/auth")
	// {
	// 	adminRoute.POST("/signin", h.AdminSignIn)
	// 	adminRoute.GET("/me", m.AuthHandle(), h.AdminGetMe)
	// 	adminRoute.GET("/refresh", m.RefreshHandle())
	// }

	// serverRoute := r.Group("/server/auth", m.BasicHandle())
	// {
	// 	serverRoute.POST("/token/validate", h.ValidateAuthToken)
	// }
}
