package main

import (
	_ "embed"
	"ev-plugin/backend/migrate"
	"ev-plugin/backend/router"
	"ev-plugin/frontend"
	"flag"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server"
)

//go:embed plugin.json
var pluginJsonBytes []byte

func main() {

	flag.Parse()

	plugin_server.Serve(plugin_server.ServeOpts{
		PluginJsonBytes: pluginJsonBytes,
		Migration:       migrate.GetMigrates(),
		FrontendFiles:   frontend.StatisFs,
		WebEngine:       router.NewRouter(),
		ExitCallback: func() {
			logger.DefaultLogger.Debug("进程退出")
		},
	})
}
