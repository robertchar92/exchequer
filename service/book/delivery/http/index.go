package http

import (
	"exchequer/service/book/delivery/http/request"
	response_util "exchequer/utils/response_utils"
	"exchequer/utils/role"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	books, bookPagination, err := h.bookUsecase.Index(request.NewBookPaginationConfig(c.Request.URL.Query()))
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	data, err := role.GetDataJSONByRole(books, role.User)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, response_util.IndexResponse{
		Data: data,
		Meta: bookPagination,
	})
}
