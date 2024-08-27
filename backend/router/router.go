package router

import (
	"eve-plugin-vue3-template/backend/api"
	"eve-plugin-vue3-template/backend/response"
	"eve-plugin-vue3-template/frontend"
	"github.com/1340691923/eve-plugin-sdk-go/backend"
	"github.com/1340691923/eve-plugin-sdk-go/backend/resource/httpadapter"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	engine                *gin.Engine
	helloWorldContoller   *api.HelloWorldController
	dbTestWorldController *api.DbTestController
}

// 依赖注入
func NewWebServer(app *gin.Engine) *WebServer {
	baseController := api.NewBaseController(response.NewResponse())
	helloWorldContoller := api.NewHelloWorldController(baseController)
	dbTestWorldController := api.NewDbTestController(baseController)
	return &WebServer{
		engine:                app,
		helloWorldContoller:   helloWorldContoller,
		dbTestWorldController: dbTestWorldController,
	}
}

// 实现web资源接口（webapi） 可用任何实现http.Handle接口的Web框架开发 我这里用gin为例
func NewResourceHandler(app *gin.Engine) backend.CallResourceHandler {

	//前端页面
	//因为前端所用技术可以进行热更新，所以可进行脱离插件控制
	app.Use(frontend.Serve("/", frontend.EmbedFolder(frontend.StatisFs, "dist")))

	//后端api
	webSvr := NewWebServer(app)

	webSvr.engine.POST("/api/HelloWorld", webSvr.helloWorldContoller.HelloAction)
	webSvr.engine.POST("/api/DbInsert", webSvr.dbTestWorldController.InsertAction)
	webSvr.engine.POST("/api/DbDelete", webSvr.dbTestWorldController.DeleteAction)
	webSvr.engine.POST("/api/DbSearch", webSvr.dbTestWorldController.SearchAction)

	return httpadapter.New(app.Handler())
}
