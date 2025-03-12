package dto

type BaseResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status: 200,
		Message: message,
		Data: data,
	}
}

func ErrorResponse(status int, message string) BaseResponse {
	return BaseResponse{
		Status: status,
		Message: message,
	}
}