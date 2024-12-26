package router

import (
	"ev-plugin/backend/api"
	"ev-plugin/backend/response"
	"github.com/1340691923/eve-plugin-sdk-go/backend/web_engine"
)

type WebServer struct {
	engine                *web_engine.WebEngine
	helloWorldContoller   *api.HelloWorldController
	dbTestWorldController *api.DbTestController
}

// 依赖注入
func NewWebServer(app *web_engine.WebEngine) *WebServer {
	baseController := api.NewBaseController(response.NewResponse())
	helloWorldContoller := api.NewHelloWorldController(baseController)
	dbTestWorldController := api.NewDbTestController(baseController)
	return &WebServer{
		engine:                app,
		helloWorldContoller:   helloWorldContoller,
		dbTestWorldController: dbTestWorldController,
	}
}

func NewRouter() *web_engine.WebEngine {

	app := web_engine.NewWebEngine()

	//后端api
	webSvr := NewWebServer(app)

	webSvr.engine.GetGinEngine().POST("/api/HelloWorld", webSvr.helloWorldContoller.HelloAction)
	webSvr.engine.GetGinEngine().POST("/api/DbInsert", webSvr.dbTestWorldController.InsertAction)
	webSvr.engine.GetGinEngine().POST("/api/DbDelete", webSvr.dbTestWorldController.DeleteAction)
	webSvr.engine.GetGinEngine().POST("/api/DbSearch", webSvr.dbTestWorldController.SearchAction)

	return app
}
