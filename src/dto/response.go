package dto

type BaseResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

type BaseResponsePagination struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func SuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status: 200,
		Message: message,
		Data: data,
	}
}

func SuccessResponseWithPagination(message string, data interface{}, limit int, page int, totalItems int, totalPages int) BaseResponsePagination {
	return BaseResponsePagination{
		Status: 200,
		Message: message,
		Data: data,
		Limit: limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

func ErrorResponse(status int, message string) BaseResponse {
	return BaseResponse{
		Status: status,
		Message: message,
	}
}