package httpapp

import (
	"go-rest-api/docs"
	"go-rest-api/internal/config"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go Rest API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func initDocs(e *echo.Echo, config config.Config) {
	var schemes []string

	if config.IsDevelopment() {
		schemes = append(schemes, "http")
	} else {
		schemes = append(schemes, "https")
	}

	docs.SwaggerInfo.Title = "Go Rest Api"
	docs.SwaggerInfo.Description = "This is a sample api."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + config.Host.Port
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = schemes

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
