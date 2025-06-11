// ElasticView插件API控制器包 - 提供HTTP接口处理的控制器层
package api

import (
	"ev-plugin/backend/response" // 导入响应处理包
)

// BaseController 基础控制器结构体 - 所有业务控制器的父类
// 采用组合模式，嵌入Response结构体，为所有子控制器提供统一的响应处理能力
type BaseController struct {
	*response.Response // 嵌入响应处理器，提供Success和Error方法
}

// NewBaseController 构造函数 - 创建基础控制器实例
// 参数 response: 响应处理器实例，用于处理HTTP响应
// 返回值: 初始化完成的BaseController实例
func NewBaseController(response *response.Response) *BaseController {
	return &BaseController{Response: response} // 注入响应处理器
}
