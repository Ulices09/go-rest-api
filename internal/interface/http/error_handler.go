package httpapp

import (
	"go-rest-api/internal/core/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if customError, ok := err.(errors.CustomError); ok {
		c.JSON(customError.Status, customError)
		return
	}

	if e, ok := err.(*echo.HTTPError); ok {
		c.JSON(e.Code, e)
		return
	}

	c.NoContent(http.StatusInternalServerError)
}
