package main

import (
	"puente/appconfig"
	"puente/cmd"
	"puente/httpserver"

	"go.uber.org/fx"
)

func main() {
	var configPath = cmd.ReadConfigPath()
	app := fx.New(
		fx.Supply(configPath),
		fx.Provide(
			appconfig.LoadAppConfig,
		),
		httpserver.Module,
		fx.Invoke(httpserver.StartWebServer),
	)
	app.Run()
}
