// ElasticView插件HelloWorld VO包 - 定义HelloWorld接口的响应数据结构
package vo

import "encoding/json" // 导入JSON包，用于处理JSON原始消息

// HelloWorld 结构体 - HelloWorld接口响应数据
// 用于向前端返回ElasticSearch连接测试的结果数据
type HelloWorld struct {
	Text json.RawMessage `json:"text"` // 原始JSON消息：保存ES ping响应的原始JSON数据，不进行二次序列化
}
