// ElasticView插件数据库测试控制器 - 演示插件数据存储功能的CRUD操作
package api

import (
	"ev-plugin/backend/dto"      // 导入数据传输对象包
	"ev-plugin/backend/model"    // 导入数据模型包
	"ev-plugin/backend/response" // 导入响应处理包

	"github.com/gin-gonic/gin" // 导入Gin HTTP框架
)

// DbTestController 数据库测试控制器结构体
// 继承BaseController，用于演示ElasticView插件的数据存储能力
type DbTestController struct {
	*BaseController // 继承基础控制器，获得统一的响应处理能力
}

// NewDbTestController 构造函数 - 创建数据库测试控制器实例
// 参数 baseController: 基础控制器实例，提供响应处理能力
// 返回值: 初始化完成的DbTestController实例
func NewDbTestController(baseController *BaseController) *DbTestController {
	return &DbTestController{BaseController: baseController} // 注入基础控制器
}

// InsertAction HTTP处理方法 - 处理数据插入请求
// 功能：向插件数据库中插入新的测试数据记录
// 参数 ctx: Gin框架的上下文对象，包含HTTP请求和响应信息
func (this *DbTestController) InsertAction(ctx *gin.Context) {
	req := new(dto.DbTestInsertReq) // 创建插入请求DTO实例
	err := ctx.Bind(req)            // 绑定请求参数到DTO结构体
	if err != nil {                 // 检查参数绑定是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	// 创建测试数据模型实例，填充从请求中获取的数据
	testModel := &model.TestModel{
		Key:   req.Key,   // 设置键值
		Value: req.Value, // 设置数据值
	}

	// 调用模型的Insert方法，执行数据库插入操作
	err = testModel.Insert()

	if err != nil { // 检查插入操作是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	// 返回操作成功响应，数据为nil表示只返回状态信息
	this.Success(ctx, response.OperateSuccess, nil)
}

// DeleteAction HTTP处理方法 - 处理数据删除请求
// 功能：清空插件数据库中的所有测试数据
// 参数 ctx: Gin框架的上下文对象，包含HTTP请求和响应信息
func (this *DbTestController) DeleteAction(ctx *gin.Context) {

	testModel := new(model.TestModel) // 创建测试数据模型实例

	// 调用模型的Clean方法，清空表中所有数据
	err := testModel.Clean()

	if err != nil { // 检查删除操作是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	// 返回操作成功响应，数据为nil表示只返回状态信息
	this.Success(ctx, response.OperateSuccess, nil)
}

// SearchAction HTTP处理方法 - 处理数据查询请求
// 功能：分页查询插件数据库中的测试数据，并返回总数统计
// 参数 ctx: Gin框架的上下文对象，包含HTTP请求和响应信息
func (this *DbTestController) SearchAction(ctx *gin.Context) {

	req := new(dto.DbTestSearchReq) // 创建查询请求DTO实例
	err := ctx.BindJSON(req)        // 绑定JSON请求体到DTO结构体
	if err != nil {                 // 检查参数绑定是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	testModel := new(model.TestModel) // 创建测试数据模型实例

	// 调用模型的List方法，执行分页查询
	// 参数：页码和每页数量
	list, err := testModel.List(req.Page, req.Limit)

	if err != nil { // 检查查询操作是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}

	// 调用模型的Count方法，获取总记录数
	count, err := testModel.Count()
	if err != nil { // 检查计数操作是否出错
		this.Error(ctx, err) // 如果出错，返回错误响应
		return
	}
	// 返回查询成功响应，包含数据列表和总数统计
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list, "count": count})
}
