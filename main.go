// 包声明：定义当前文件所属的包名
package main

import (
	"context"
	"embed"
	_ "embed"                   // 匿名导入，用于支持embed功能
	"ev-plugin/backend/migrate" // 导入数据库迁移相关模块
	"ev-plugin/backend/router"  // 导入路由相关模块
	"ev-plugin/frontend"        // 导入前端相关模块
	"flag"                      // 导入命令行参数解析包

	"github.com/1340691923/eve-plugin-sdk-go/backend/plugin_server" // 导入Eve插件服务器SDK
	"github.com/1340691923/eve-plugin-sdk-go/build"                 // 导入构建相关工具
)

// 嵌入插件配置文件，将plugin.json文件内容嵌入到二进制文件中
//
//go:embed plugin.json
var pluginJsonBytes []byte

// 嵌入插件图标文件，将logo.png文件嵌入到二进制文件中
//
//go:embed logo.png
var logoPng embed.FS

// main函数：程序入口点
func main() {

	// 解析命令行参数
	flag.Parse()

	// 启动插件服务器
	plugin_server.Serve(plugin_server.ServeOpts{
		// 配置插件资源
		Assets: &plugin_server.Assets{
			PluginJsonBytes: pluginJsonBytes,   // 插件配置文件字节数组
			FrontendFiles:   frontend.StatisFs, // 前端静态文件系统
			Icon:            logoPng,           // 插件图标
		},
		// 服务器就绪后的回调函数
		ReadyCallBack: func(ctx context.Context) {
			// 这里可以添加插件启动后需要执行的初始化逻辑
		},
		// 数据库迁移配置
		Migration: &build.Gormigrate{Migrations: []*build.Migration{
			migrate.V0_0_1(), // 执行版本0.0.1的数据库迁移
		}}, // 数据版本迁移管理
		// 注册路由处理器
		RegisterRoutes: router.NewRouter, // 注册插件的http接口路由配置
	})
}
