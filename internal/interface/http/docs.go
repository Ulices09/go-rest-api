package httpapp

import (
	"go-rest-api/docs"
	"go-rest-api/internal/config"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @securityDefinitions.apikey auth-token
// @in cookie
// @name session-token
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
