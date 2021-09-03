package main

import (
	"go-rest-api/db"
	"go-rest-api/domain/posts"

	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	db := db.InitDb()
	defer db.Close()

	postRepo := posts.NewPostRepository(db)
	postService := posts.NewPostService(postRepo)
	postController := posts.NewPostController(postService)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, uri=${uri}, status=${status}\n",
	}))

	posts.NewPostRouter(e, postController)

	e.Logger.Fatal(e.Start((":8000")))
}
