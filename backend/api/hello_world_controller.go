// ElasticView插件HelloWorld控制器 - 演示插件与ElasticSearch连接测试功能
package api

import (
	"ev-plugin/backend/dto"      // 导入数据传输对象包
	"ev-plugin/backend/response" // 导入响应处理包
	"ev-plugin/backend/vo"       // 导入视图对象包

	"github.com/1340691923/eve-plugin-sdk-go/backend/logger" // 导入ElasticView插件SDK日志组件
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"         // 导入ElasticView插件SDK API组件
	"github.com/1340691923/eve-plugin-sdk-go/util"           // 导入ElasticView插件SDK工具包
	"github.com/gin-gonic/gin"                               // 导入Gin HTTP框架
)

// HelloWorldController HelloWorld功能控制器结构体
// 继承BaseController，具备统一的响应处理能力
type HelloWorldController struct {
	*BaseController // 继承基础控制器
}

// NewHelloWorldController 构造函数 - 创建HelloWorld控制器实例
// 参数 baseController: 基础控制器实例，提供响应处理能力
// 返回值: 初始化完成的HelloWorldController实例
func NewHelloWorldController(baseController *BaseController) *HelloWorldController {
	return &HelloWorldController{baseController} // 注入基础控制器
}

// HelloAction HTTP处理方法 - 处理HelloWorld接口请求
// 功能：测试插件与指定ElasticSearch连接的连通性
// 参数 ctx: Gin框架的上下文对象，包含HTTP请求和响应信息
func (this *HelloWorldController) HelloAction(ctx *gin.Context) {
	req := new(dto.HelloWorld) // 创建请求DTO实例
	err := ctx.BindJSON(req)   // 绑定JSON请求体到DTO结构体
	if err != nil {            // 检查参数绑定是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	// 使用ElasticView插件SDK的日志组件记录调试信息
	logger.DefaultLogger.Debug("打印日志", "conn_id:", req.EsConnect)

	// 创建ElasticView API包装器实例
	// 参数1: ES连接ID，通常由前端传递（sdk.GetSelectEsConnID()）
	// 参数2: 当前用户ID，从上下文中获取
	esI := ev_api.NewEvWrapApi(req.EsConnect, util.GetEvUserID(ctx))

	// 执行ElasticSearch连接测试
	// 调用Ping方法检查指定ES连接的连通性
	res, err := esI.Ping(ctx)

	if err != nil { // 检查Ping操作是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	// 返回成功响应
	// 参数1: Gin上下文
	// 参数2: 成功消息常量
	// 参数3: 响应数据VO，包含ES连接测试的原始JSON结果
	this.Success(ctx, response.SearchSuccess, vo.HelloWorld{
		Text: res.JsonRawMessage(), // 将ES响应转换为JSON原始消息
	})
}
