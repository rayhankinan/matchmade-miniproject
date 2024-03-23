package types

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type PaginatedResponse struct {
	Data         interface{} `json:"data"`
	TotalPages   int64       `json:"total_pages"`
	TotalResults int64       `json:"total_results"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
