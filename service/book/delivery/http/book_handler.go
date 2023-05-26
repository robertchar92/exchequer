package http

import (
	"exchequer/app/middleware"
	"exchequer/service/book"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	bookUsecase book.Usecase
}

func New(bookUC book.Usecase) *Handler {
	return &Handler{
		bookUsecase: bookUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	userRoute := r.Group("/book", m.AuthHandle(), m.UserHandle())
	{
		userRoute.GET("", h.Index)
		userRoute.GET("/:id", h.Show)
		userRoute.POST("", h.Create)
		userRoute.PATCH("/:id", h.Update)
	}
}
