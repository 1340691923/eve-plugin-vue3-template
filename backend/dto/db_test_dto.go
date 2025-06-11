// ElasticView插件数据库测试DTO包 - 定义数据库操作接口的请求参数结构
package dto

// DbTestSearchReq 结构体 - 数据库查询请求参数
// 用于接收前端传递的分页查询参数
type DbTestSearchReq struct {
	Page  int `json:"page"`  // 页码：指定查询第几页，从1开始计数
	Limit int `json:"limit"` // 每页数量：指定每页返回的记录数量
}

// DbTestInsertReq 结构体 - 数据库插入请求参数
// 用于接收前端传递的数据插入参数
type DbTestInsertReq struct {
	Key   string `json:"key"`   // 键名：要插入的数据键名
	Value string `json:"value"` // 键值：要插入的数据值
}
