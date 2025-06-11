// ElasticView插件自定义异常处理包 - 提供统一的错误处理机制
package my_error

// MyError 自定义异常结构体 - 实现Go标准库的error接口
// 用于在ElasticView插件中统一处理和管理错误信息
type MyError struct {
	code int    // 错误码 - 用于标识具体的错误类型
	msg  string // 错误消息 - 用于描述错误的详细信息
}

// NewError 构造函数 - 创建新的自定义错误实例
// 参数 text: 错误描述信息
// 参数 code: 错误码，通常0表示成功，非0表示异常
// 返回值: 实现了error接口的MyError实例
func NewError(text string, code int) error {
	return &MyError{code, text} // 返回MyError指针，实现error接口
}

// Error 方法 - 实现Go标准库error接口的Error()方法
// 返回值: 错误的字符串描述
func (this *MyError) Error() string {
	return this.msg // 返回错误消息
}

// Code 方法 - 获取错误码的自定义方法
// 返回值: 错误码整数值
func (this *MyError) Code() int {
	return this.code // 返回错误码
}
