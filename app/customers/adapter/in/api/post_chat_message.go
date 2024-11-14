package api

import (
	"einar-website-api/app/shared/infrastructure/serverwrapper"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(postChatMessage, serverwrapper.NewEchoWrapper)
}
func postChatMessage(e serverwrapper.EchoWrapper) {
	e.POST("/chat-message", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Pronto estaremos disponibles 🚀.",
		})
	})
}
