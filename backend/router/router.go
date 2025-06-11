// ElasticView插件路由配置包 - 负责注册插件的HTTP接口路由
package router

import (
	"ev-plugin/backend/api"      // 导入API控制器包
	"ev-plugin/backend/response" // 导入响应处理包

	"github.com/1340691923/eve-plugin-sdk-go/backend/web_engine" // 导入ElasticView插件SDK的Web引擎
)

// WebServer 结构体 - 插件Web服务器的核心结构
// 包含Web引擎和各种控制器的实例
type WebServer struct {
	engine                *web_engine.WebEngine     // ElasticView插件SDK提供的Web引擎实例
	helloWorldContoller   *api.HelloWorldController // HelloWorld功能控制器
	dbTestWorldController *api.DbTestController     // 数据库测试功能控制器
}

// NewWebServer 构造函数 - 使用依赖注入模式创建WebServer实例
// 参数 app: ElasticView插件SDK提供的Web引擎实例
// 返回值: 初始化完成的WebServer实例
func NewWebServer(app *web_engine.WebEngine) *WebServer {
	baseController := api.NewBaseController(response.NewResponse())    // 创建基础控制器，注入响应处理器
	helloWorldContoller := api.NewHelloWorldController(baseController) // 创建HelloWorld控制器，注入基础控制器
	dbTestWorldController := api.NewDbTestController(baseController)   // 创建数据库测试控制器，注入基础控制器
	return &WebServer{
		engine:                app,                   // 注入Web引擎
		helloWorldContoller:   helloWorldContoller,   // 注入HelloWorld控制器
		dbTestWorldController: dbTestWorldController, // 注入数据库测试控制器
	}
}

// NewRouter 函数 - ElasticView插件路由注册的入口函数
// 这个函数会被主程序调用，用于注册插件的所有HTTP接口路由
// 参数 engine: ElasticView插件SDK提供的Web引擎实例
func NewRouter(engine *web_engine.WebEngine) {

	// 创建WebServer实例，完成依赖注入
	webSvr := NewWebServer(engine)

	// 注册HelloWorld接口 - POST方法，路径/api/HelloWorld，处理器为HelloAction
	webSvr.engine.GetGinEngine().POST("/api/HelloWorld", webSvr.helloWorldContoller.HelloAction)
	// 注册数据库插入接口 - POST方法，路径/api/DbInsert，处理器为InsertAction
	webSvr.engine.GetGinEngine().POST("/api/DbInsert", webSvr.dbTestWorldController.InsertAction)
	// 注册数据库删除接口 - POST方法，路径/api/DbDelete，处理器为DeleteAction
	webSvr.engine.GetGinEngine().POST("/api/DbDelete", webSvr.dbTestWorldController.DeleteAction)
	// 注册数据库查询接口 - POST方法，路径/api/DbSearch，处理器为SearchAction
	webSvr.engine.GetGinEngine().POST("/api/DbSearch", webSvr.dbTestWorldController.SearchAction)

}
