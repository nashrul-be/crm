package dto

import "net/http"

type BaseResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func ErrorNotFound(msgErr string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusNotFound,
		Message: msgErr,
	}
}

func ErrorBadRequest(msgErr string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusBadRequest,
		Message: msgErr,
	}
}

func ErrorInternalServerError() BaseResponse {
	return BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: "Oops, something wrong!",
	}
}

func ErrorUnauthorized(msg string) BaseResponse {
	return BaseResponse{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func Success(msg string, data any) BaseResponse {
	return BaseResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func Created(msg string, data any) BaseResponse {
	return BaseResponse{
		Code:    http.StatusCreated,
		Message: msg,
		Data:    data,
	}
}
