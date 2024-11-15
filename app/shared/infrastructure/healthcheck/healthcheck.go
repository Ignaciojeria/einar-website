package api

import (
	"einar-website-api/app/shared/configuration"
	"einar-website-api/app/shared/infrastructure/serverwrapper"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(healthCheck,
		serverwrapper.NewEchoWrapper,
		configuration.NewConf)
}

// To see usage examples of the library, visit: https://github.com/hellofresh/health-go
func healthCheck(e serverwrapper.EchoWrapper, c configuration.Conf) {
	h, _ := health.New(
		health.WithComponent(health.Component{
			Name:    c.PROJECT_NAME,
			Version: c.VERSION,
		}), health.WithSystemInfo())
	e.GET("/health", echo.WrapHandler(h.Handler()))
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "")
	})
}
