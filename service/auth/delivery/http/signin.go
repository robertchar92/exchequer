package http

import (
	"exchequer/service/auth/delivery/http/request"
	"exchequer/utils/role"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var req request.SignInRequest

	// validate request
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	res, err := h.authUsecase.SignIn(req)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	data, err := role.GetDataJSONByRole(res, role.User)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, data)
}
