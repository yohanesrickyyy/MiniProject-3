package helper

import (
	"miniproject-3/model/web"
	"net/http"
)

func ErrBadRequest(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
		Detail:  detail,
	}
}

func ErrUnauthorized(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized access",
		Detail:  detail,
	}
}

func ErrInternalServer(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Detail:  detail,
	}
}

func ErrNotFound(detail any) web.ErrorResponse {
	return web.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Error not found",
		Detail:  detail,
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
