package topbar

import (
	"embed"
	"my-project-name/app/adapter/in/view/component"
	"my-project-name/app/shared/archetype/container"
	einar "my-project-name/app/shared/archetype/echo_server"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:embed *.html
var html embed.FS

func init() {
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: html,
		Pattern: component.Topbar + component.DOT_HTML,
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/"+component.Topbar, render)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func render(c echo.Context) error {
	routerState := einar.NewRoutingState(c, map[string]string{
		component.IndexComponentDefault:  component.Topbar,
		component.TopbarComponentDefault: component.Home,
	})
	if c.Request().Header.Get(component.FlatContext) != "" {
		return c.Render(http.StatusOK, component.Topbar+component.DOT_HTML, routerState)
	}
	return c.Render(http.StatusOK, component.Index+component.DOT_HTML, routerState)
}
