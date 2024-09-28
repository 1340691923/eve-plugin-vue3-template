package main

import (
	_ "embed"
	"encoding/json"
	"eve-plugin-vue3-template/backend/global"
	"eve-plugin-vue3-template/backend/migrate"
	"eve-plugin-vue3-template/backend/router"
	"log"

	"flag"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server"
	"github.com/1340691923/eve-plugin-sdk-go/build"
	"github.com/gin-gonic/gin"
)

//go:embed plugin.json
var pluginJsonBytes []byte

func init() {
	flag.StringVar(&global.EvRpcPort, "evRpcPort", "", "ev基座内网访问端口")
	flag.StringVar(&global.EvRpcKey, "evRpcKey", "", "ev基座访问密钥")
	flag.StringVar(&global.TmpFileStorePath, "tmpFileStorePath", "store_file_dir", "临时文件存放目录")
	flag.BoolVar(&global.Debug, "debug", false, "是否开启调试")
	flag.Parse()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	pluginJson := new(build.PluginJsonData)
	err := json.Unmarshal(pluginJsonBytes, &pluginJson)
	if err != nil {
		log.Println("plugin.json解析失败")
		panic(err)
	}

	pluginJson.BackendDebug = global.Debug
	plugin_server.Serve(plugin_server.ServeOpts{
		EvRpcPort:           global.EvRpcPort,
		EvRpcKey:            global.EvRpcKey,
		PluginJson:          pluginJson,
		Migration:           migrate.GetMigrates(),
		CallResourceHandler: router.NewResourceHandler(gin.New()),
		ExitCallback: func() {
			logger.DefaultLogger.Debug("进程退出")
		},
	})
}
