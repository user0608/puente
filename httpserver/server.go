package httpserver

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"puente/appconfig"
	"puente/httpserver/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module("http-server",
	fx.Provide(
		fx.Annotate(
			newEchoServer,
			fx.ParamTags(
				routes.RouteTag,
			),
		),
	),
)

func newEchoServer(listRoutes []routes.Route) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_unix}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))
	e.Use(middleware.Recover())
	for _, route := range listRoutes {
		method, path := route.MethodAndPath()
		var middlewares = []echo.MiddlewareFunc{}
		if route, ok := route.(routes.RouteWithMiddles); ok {
			middlewares = route.Use()
		}
		switch method {
		case http.MethodGet:
			e.GET(path, route.Handle, middlewares...)
		case http.MethodPost:
			e.POST(path, route.Handle, middlewares...)
		case http.MethodPut:
			e.PUT(path, route.Handle, middlewares...)
		case http.MethodDelete:
			e.DELETE(path, route.Handle, middlewares...)
		case http.MethodPatch:
			e.PATCH(path, route.Handle, middlewares...)
		default:
			slog.Warn("unsuported method was found", "method", method)
		}
	}
	return e
}

func StartWebServer(lc fx.Lifecycle, e *echo.Echo, c appconfig.AppConfig) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			e.HideBanner = true
			go func() {
				var err = e.Start(c.ListenAddress())
				if err != nil && err != http.ErrServerClosed {
					log.Println(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
