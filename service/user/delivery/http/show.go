package http

import (
	"exchequer/utils/role"
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Show(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	res, err := h.userUsecase.Show(fmt.Sprint(claims["id"]), true)
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
