package http

import (
	"exchequer/service/auth/delivery/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req request.SignUpRequest

	// validate request
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	err := h.authUsecase.SignUp(req)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
