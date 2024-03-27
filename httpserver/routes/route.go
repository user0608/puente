package routes

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const RouteTag = `group:"http-routes"`

// Route defines the structure for an HTTP route.
type Route interface {
	MethodAndPath() (method string, path string)
	Handle(c echo.Context) error
}

type RouteWithMiddles interface {
	Use() []echo.MiddlewareFunc
	Route
}

func AsRoute(fn any) any {
	return fx.Annotate(
		fn,
		fx.As(new(Route)),
		fx.ResultTags(RouteTag),
	)
}
