package book

import (
	"exchequer/models"
	"exchequer/service/book/delivery/http/request"
	"exchequer/utils/request_util"
	response_util "exchequer/utils/response_utils"
)

type Usecase interface {
	Create(userID string, username string, request request.BookCreateRequest) (*models.Book, error)
	Update(bookID string, userID string, request request.BookUpdateRequest) (*models.Book, error)
	Index(paginationConfig request_util.PaginationConfig) ([]models.Book, response_util.PaginationMeta, error)
	Show(bookID string, userID string) (*models.Book, error)
}
