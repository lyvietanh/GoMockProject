package util

import (
	"fmt"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateSuccessResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Success: false,
		Code:    200,
		Message: "",
		Data:    data,
	}
}

// func CreateErrorResponse(err error) ApiResponse {
// 	return ApiResponse{
// 		Success: false,
// 		Code:    500,
// 		Message: err.Error(),
// 		Data:    nil,
// 	}
// }

// func CreateCustomErrorResponse(errorCode int, errorMessage string, data interface{}) ApiResponse {
// 	return ApiResponse{
// 		Success: false,
// 		Code:    errorCode,
// 		Message: errorMessage,
// 		Data:    data,
// 	}
// }

func CreateApiErrorResponse(apiError ApiError, data interface{}) ApiResponse {
	return ApiResponse{
		Success: false,
		Code:    apiError.Code,
		Message: apiError.Message,
		Data:    data,
	}
}

func CreateCustomApiErrorResponse(apiError ApiError, additionalMessage string, data interface{}) ApiResponse {
	return ApiResponse{
		Success: false,
		Code:    apiError.Code,
		Message: fmt.Sprintf("%v|%v", apiError.Message, additionalMessage),
		Data:    data,
	}
}
