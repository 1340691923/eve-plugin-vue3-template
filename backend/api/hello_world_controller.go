package api

import (
	"eve-plugin-vue3-template/backend/dto"
	"eve-plugin-vue3-template/backend/response"
	"eve-plugin-vue3-template/backend/vo"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// hello world控制器
type HelloWorldController struct {
	*BaseController
}

func NewHelloWorldController(baseController *BaseController) *HelloWorldController {
	return &HelloWorldController{baseController}
}

func (this *HelloWorldController) HelloAction(ctx *gin.Context) {
	req := new(dto.HelloWorld)
	err := ctx.BindJSON(req)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	logger.DefaultLogger.Debug("打印日志", "conn_id:", req.EsConnect)

	//调用基座api
	esI := ev_api.NewEvWrapApi(req.EsConnect, cast.ToInt(ctx.GetHeader(util.EvRoleID)), cast.ToInt(ctx.GetHeader(util.EvUserID)))

	//ping一下所在连接的es
	res, err := esI.Ping(ctx)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, vo.HelloWorld{
		Text: res.JsonRawMessage(),
	})
}
