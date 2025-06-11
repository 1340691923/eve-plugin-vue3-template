// ElasticView插件HelloWorld DTO包 - 定义HelloWorld接口的请求参数结构
package dto

// HelloWorld 结构体 - HelloWorld接口请求参数
// 用于接收前端传递的ElasticSearch连接测试参数
type HelloWorld struct {
	EsConnect int    `json:"es_connect"` // ES连接ID：指定要测试的ElasticSearch连接标识符
	Text      string `json:"text"`       // 文本内容：可选的文本参数，用于扩展功能
}
