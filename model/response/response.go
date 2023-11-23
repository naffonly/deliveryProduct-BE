package response

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
}

func FormatPaginationResponse(message string, data any, pagination any) map[string]any {
	var responsePagination = map[string]any{}

	responsePagination["message"] = message
	responsePagination["data"] = data
	responsePagination["pagination"] = pagination

	return responsePagination
}
