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
	// userRoute := r.Group("/auth")
	// {
	// versionCheckGroup := userRoute.Group("")
	// {
	// 	versionCheckGroup.POST("/signin", m.VersionHandle(), h.SignInWithUsername)
	// 	versionCheckGroup.POST("/signin/phone", m.VersionHandle(), h.SignInWithPhone)
	// 	versionCheckGroup.POST("/signin/phone/validate", m.VersionHandle(), h.SignInWithPhoneValidate)
	// 	versionCheckGroup.POST("/phone", m.VersionHandle(), h.SignUpWithPhone)
	// 	versionCheckGroup.POST("/phone/validate", m.VersionHandle(), h.SignUpWithPhoneValidate)
	// 	versionCheckGroup.POST("/signup", m.VersionHandle(), h.SignUp)
	// 	versionCheckGroup.POST("/otp/phone", m.VersionHandle(), h.RequestPhoneOTP)
	// 	versionCheckGroup.POST("/otp/phone/validate", m.VersionHandle(), h.ValidatePhoneOTP)
	// }
	// userRoute.POST("/forget-password", h.ForgetPassword)
	// userRoute.GET("/reset-password/:token", h.ValidateResetPasswordToken)
	// userRoute.POST("/reset-password", h.ResetPassword)

	// refresh jwt token
	// userRoute.GET("/refresh", m.RefreshHandle())
	// userRoute.GET("/temp-token", m.AuthHandle(), h.TempToken)
	// userRoute.POST("/google-callback", h.GoogleCallback)
	// userRoute.POST("/facebook-callback", h.FacebookCallback)
	// }

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
