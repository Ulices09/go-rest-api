package main

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

type ErrorField struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		errorFields := []ErrorField{}

		for _, err := range err.(validator.ValidationErrors) {

			fields := strings.Split(err.Namespace(), ".")
			errField := ""

			for i, f := range fields {

				if i == 0 {
					continue
				}

				errField += strings.ToLower(f)

				if i+1 != len(fields) {
					errField += "."
				}

			}

			errorField := ErrorField{Field: errField, Code: err.Tag(), Message: err.Error()}
			errorFields = append(errorFields, errorField)

		}

		return echo.NewHTTPError(http.StatusBadRequest, errorFields)
	}
	return nil
}

func main() {
	app := echo.New()

	app.Validator = &CustomValidator{validator: validator.New()}

	app.GET("", GetPosts)
	app.GET(":id", GetPost)
	app.POST("", CreatePost)

	app.Logger.Fatal(app.Start((":8000")))
}
