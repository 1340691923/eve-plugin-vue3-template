// ElasticView插件自定义响应处理包 - 提供统一的HTTP响应格式和错误处理机制
package response

import (
	"encoding/json"              // 导入JSON编码包
	"ev-plugin/backend/my_error" // 导入自定义错误处理包
	"fmt"                        // 导入格式化包，用于字符串格式化
	"io"                         // 导入IO包，用于数据流操作
	"net/http"                   // 导入HTTP包，用于HTTP状态码

	"github.com/1340691923/eve-plugin-sdk-go/backend/logger" // 导入ElasticView插件SDK日志组件
	"github.com/gin-gonic/gin"                               // 导入Gin HTTP框架
)

// ResponseData 结构体 - 定义统一的HTTP响应数据格式
// ElasticView插件采用统一的响应格式，便于前端统一处理
type ResponseData struct {
	Code int         `json:"code"` // 响应状态码：0表示成功，非0表示失败
	Msg  string      `json:"msg"`  // 响应消息：提供给用户的提示信息
	Data interface{} `json:"data"` // 响应数据：实际的业务数据，可以是任意类型
}

// Response 结构体 - 自定义响应处理器
// 提供统一的成功和错误响应方法
type Response struct {
}

// NewResponse 构造函数 - 创建响应处理器实例
// 返回值: Response实例指针
func NewResponse() *Response {
	return &Response{} // 返回空的Response实例
}

// 响应状态码常量定义
const (
	SUCCESS = 0   // 成功状态码
	ERROR   = 500 // 错误状态码，对应HTTP 500内部服务器错误
)

// 响应消息常量定义 - 统一管理常用的响应消息
const (
	SearchSuccess       = "查询成功"   // 查询操作成功消息
	DeleteSuccess       = "删除成功"   // 删除操作成功消息
	OperateSuccess      = "操作成功"   // 通用操作成功消息
	LogoutSuccess       = "注销成功"   // 注销操作成功消息
	ChangeLayoutSuccess = "修改布局成功" // 布局修改成功消息
)

// Success 方法 - 处理成功响应
// 统一的成功响应格式，返回成功状态码和业务数据
// 参数 ctx: Gin上下文对象
// 参数 msg: 成功消息字符串
// 参数 data: 业务数据，可以是任意类型
// 返回值: 错误信息，始终返回nil
func (this *Response) Success(ctx *gin.Context, msg string, data interface{}) error {
	responseData := new(ResponseData) // 创建响应数据实例
	responseData.Msg = msg            // 设置响应消息
	responseData.Data = data          // 设置业务数据
	responseData.send(ctx, SUCCESS)   // 发送成功响应
	return nil                        // 返回nil表示处理成功
}

// Error 方法 - 处理错误响应
// 统一的错误响应格式，记录错误日志并返回错误信息给前端
// 参数 ctx: Gin上下文对象
// 参数 err: 错误信息
// 返回值: 错误信息，始终返回nil（错误已被处理）
func (this *Response) Error(ctx *gin.Context, err error) error {

	myErr := ErrorToErrorCode(err) // 将标准错误转换为自定义错误格式

	var b []byte            // 声明字节切片变量
	b, _ = ctx.GetRawData() // 获取原始请求体数据

	// 使用ElasticView插件SDK的日志组件记录错误信息
	logger.DefaultLogger.Error("错误日志",
		"请求接口地址:", ctx.Request.URL.Path, // 记录请求的接口地址
		"请求body:", string(b), // 记录请求体内容
		"异常堆栈:", fmt.Sprintf("%+v", err), // 记录详细的错误堆栈信息
	)

	responseData := new(ResponseData)    // 创建响应数据实例
	responseData.Msg = myErr.Error()     // 设置错误消息
	responseData.send(ctx, myErr.Code()) // 发送错误响应
	return nil                           // 返回nil表示错误已被处理
}

// send 方法 - 发送HTTP响应
// 根据状态码决定HTTP响应的状态码，并将响应数据序列化为JSON返回
// 参数 ctx: Gin上下文对象
// 参数 code: 业务状态码
// 返回值: 错误信息，始终返回nil
func (this *ResponseData) send(ctx *gin.Context, code int) error {
	this.Code = code    // 设置响应状态码
	if this.Code != 0 { // 如果不是成功状态码
		ctx.JSON(http.StatusAccepted, this) // 返回HTTP 202（已接受）状态，表示请求已处理但有业务错误
	} else { // 如果是成功状态码
		ctx.JSON(http.StatusOK, this) // 返回HTTP 200（成功）状态
	}
	return nil // 返回nil表示发送成功
}

// Output 方法 - 通用数据输出方法
// 将数据序列化为JSON格式并写入到指定的Writer中
// 参数 write: 数据写入目标
// 参数 data: 要输出的数据
// 返回值: 错误信息
func (this *Response) Output(write io.Writer, data map[string]interface{}) error {
	b, _ := json.Marshal(data) // 将数据序列化为JSON字节数组
	write.Write(b)             // 将JSON数据写入Writer
	return nil                 // 返回nil表示输出成功
}

// ErrorToErrorCode 函数 - 错误类型转换工具函数
// 将Go标准错误类型转换为ElasticView插件的自定义错误类型
// 参数 err: 标准错误接口
// 返回值: 自定义错误类型指针
func ErrorToErrorCode(err error) *my_error.MyError {
	if err == nil { // 如果错误为空
		return nil // 返回nil
	}

	// 尝试将错误断言为自定义错误类型
	errorCode, ok := err.(*my_error.MyError)

	if ok { // 如果断言成功
		return errorCode // 直接返回自定义错误
	}
	// 如果断言失败，创建新的自定义错误
	return my_error.NewError(err.Error(), ERROR).(*my_error.MyError)
}
