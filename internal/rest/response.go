package rest

import (
	"github.com/labstack/echo/v4"
)

func newErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.String(statusCode, message)
}

var serverError = `server error`
