package http

import "github.com/labstack/echo/v4"

func NewError(ctx echo.Context, status int, err error) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Bad request"`
}
