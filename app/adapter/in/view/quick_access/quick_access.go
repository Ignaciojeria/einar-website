package quick_access

import (
	"my-project-name/app/adapter/in/view/component"
	"my-project-name/app/shared/archetype/container"
	einar "my-project-name/app/shared/archetype/echo_server"
	"embed"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:embed *.html
var html embed.FS

//go:embed *.css
var css embed.FS

func init() {
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: html,
		Pattern: component.QuickAccess + component.DOT_HTML,
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: component.QuickAccess + component.DOT_CSS,
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/"+component.QuickAccess, render)
		einar.Echo.GET("/"+component.QuickAccess+component.DOT_CSS, echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func render(c echo.Context) error {
	routerState := einar.NewRoutingState(c, map[string]string{
		component.IndexComponentDefault: component.QuickAccess,
	})
	if c.Request().Header.Get(component.FlatContext) != "" {
		return c.Render(http.StatusOK, component.QuickAccess+component.DOT_HTML, routerState)
	}
	return c.Render(http.StatusOK, component.Index+component.DOT_HTML, routerState)
}
