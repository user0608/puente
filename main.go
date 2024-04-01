package main

import (
	"puente/appconfig"
	"puente/cmd"
	"puente/connection"
	"puente/httpserver"
	"puente/migrate"

	"go.uber.org/fx"
)

func main() {
	var configPath = cmd.ReadConfigPath()
	app := fx.New(
		fx.Supply(configPath),
		fx.Provide(appconfig.LoadAppConfig),
		fx.Provide(connection.NewConnection),
		httpserver.Module,
		fx.Invoke(migrate.ExecMigration),
		fx.Invoke(httpserver.StartWebServer),
	)
	app.Run()
}
