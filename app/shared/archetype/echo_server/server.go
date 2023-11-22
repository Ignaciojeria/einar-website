package echo_server

import (
	"fmt"
	"my-project-name/app/shared/archetype/container"
	"my-project-name/app/shared/config"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

var Echo *echo.Echo

func init() {
	config.Installations.EnableHTTPServer = true

	container.InjectInstallation(func() error {
		Echo = echo.New()
		Echo.Use(middleware.Logger())
		Echo.Use(middleware.Recover())
		// Add CORS middleware to allow all origins
		Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
		return nil
	}, container.InjectionProps{DependencyID: uuid.NewString()})

	container.InjectHTTPServer(func() error {
		setUpRenderer(EmbeddedPatterns...)
		for _, route := range Echo.Routes() {
			fmt.Printf("Method: %v, Path: %v, Name: %v\n", route.Method, route.Path, route.Name)
		}
		PORT := config.PORT.Get()
		if config.PORT.Get() == "" {
			PORT = "80"
		}
		err := Echo.Start(":" + PORT)
		if err != nil {
			log.Error().Err(err).Msg("error initializing application server")
			return err
		}
		return nil
	}, container.InjectionProps{DependencyID: uuid.NewString()})

}

func init() {
	container.InjectInboundAdapter(func() error {
		Echo.GET("/health", func(c echo.Context) error {
			return c.String(http.StatusOK, "UP")
		})
		return nil
	}, container.InjectionProps{DependencyID: uuid.NewString()})
}
