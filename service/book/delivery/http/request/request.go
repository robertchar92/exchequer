package request

import "exchequer/utils/request_util"

type BookCreateRequest struct {
	Name *string `form:"name" json:"name" binding:"required"`
}
type BookUpdateRequest struct {
	Name *string `form:"name" json:"name" binding:"omitempty"`
}

func NewBookPaginationConfig(conditions map[string][]string) request_util.PaginationConfig {
	request_util.OverrideKey(conditions, "type", "scope")

	filterable := map[string]string{
		"id":         request_util.ExactStringType,
		"user_id":    request_util.ExactStringType,
		"username":   request_util.StringType,
		"name":       request_util.StringType,
		"created_at": request_util.DateType,
	}

	conditions["sort"] = []string{"created_at DESC"}

	return request_util.NewRequestPaginationConfig(conditions, filterable)
}
