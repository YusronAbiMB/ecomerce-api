package helper

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(message string, data any) BaseResponse {
	return BaseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
func ResponseFailed(message string) BaseResponse {
	return BaseResponse{
		Status:  false,
		Message: message,
	}
}
