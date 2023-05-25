package http

import (
	"exchequer/app/middleware"
	"exchequer/service/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase user.Usecase
}

func New(userUC user.Usecase) *Handler {
	return &Handler{
		userUsecase: userUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	// adminRoute := r.Group("/admin/user", m.AuthHandle(), m.AdminHandle())
	// {
	// 	adminRoute.GET("", h.AdminIndex)
	// 	adminRoute.GET("/:id", h.AdminShow)

	// }
	userRoute := r.Group("/user", m.AuthHandle(), m.UserHandle())
	{
		userRoute.GET("", h.Show)
		userRoute.PATCH("", h.Update)
		// userRoute.POST("/email/verify", h.VerifyEmail)
		// userRoute.DELETE("", h.UserDelete)
	}

	// //public route
	// r.POST("/user/email/validate", h.ValidateEmailToken)

}
