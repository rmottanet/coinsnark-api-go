package models

import (
	"strconv"
)


type ErrorResponse struct {
	ApiResponse     
	Error           map[string]string `json:"error"`
}


func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	errorResponse := &ErrorResponse{
		Error: map[string]string{
			"message": message,
			"status":  strconv.Itoa(statusCode),
		},
	}
	errorResponse.ApiResponse = *NewApiResponse()

	return errorResponse
}
