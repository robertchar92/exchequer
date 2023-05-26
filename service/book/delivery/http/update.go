package http

import (
	"exchequer/service/book/delivery/http/request"
	"exchequer/utils/role"
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	var req request.BookUpdateRequest
	// validate request
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	res, err := h.bookUsecase.Update(c.Param("id"), fmt.Sprint(claims["id"]), req)
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
