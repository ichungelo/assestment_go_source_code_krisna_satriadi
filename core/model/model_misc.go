package model

// Presenter metadata struct
type ResponseMetadata struct {
	Success    *bool               `json:"success"`
	Code       *string             `json:"Code"`
	Error      *string             `json:"-"`
	Message    *string             `json:"message"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
}

// Presenter pagination struct
type ResponsePagination struct {
	Total *int `json:"total"`
	Count *int `json:"count"`
	Start *int `json:"start"`
}

// Response Presenter struct
type Response struct {
	Metadata ResponseMetadata `json:"metadata"`
	Result   interface{}      `json:"result"`
}
