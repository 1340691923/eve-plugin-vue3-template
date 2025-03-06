package main

import (
	"context"
	_ "embed"
	"ev-plugin/backend/migrate"
	"ev-plugin/backend/router"
	"ev-plugin/frontend"
	"flag"
	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server"
	"github.com/1340691923/eve-plugin-sdk-go/build"
	"github.com/1340691923/eve-plugin-sdk-go/live"
	"log"
)

//go:embed plugin.json
var pluginJsonBytes []byte

func main() {

	flag.Parse()

	plugin_server.Serve(plugin_server.ServeOpts{
		PluginJsonBytes: pluginJsonBytes,
		Migration: &build.Gormigrate{Migrations: []*build.Migration{
			migrate.V0_0_1(),
		}}, //数据版本迁移
		LiveHandler: live.NewLive(map[string]live.LiveChannel{
			"live": new(Live),
		}),
		FrontendFiles: frontend.StatisFs,
		WebEngine:     router.NewRouter(),
	})
}

type Live struct {
}

func (l Live) Pub2Channel(ctx context.Context, req []byte) (res map[string]interface{}, err error) {
	log.Println(string(req))
	return map[string]interface{}{"test": 123}, nil
}
