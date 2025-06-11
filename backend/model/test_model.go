// ElasticView插件数据库实体层 - 定义数据模型和数据库操作方法
package model

import (
	"context" // 导入上下文包，用于传递请求上下文

	"github.com/1340691923/eve-plugin-sdk-go/ev_api"      // 导入ElasticView插件SDK API组件
	"github.com/1340691923/eve-plugin-sdk-go/sql_builder" // 导入ElasticView插件SDK SQL构建器
)

/*
TestModel 测试数据模型结构体
ElasticView插件可以使用 http://sql2struct.atotoa.com 根据表结构生成对应的Go结构体
该模型用于演示插件的数据存储功能，记录简单的键值对数据

数据库表结构说明：
- id: 主键，自增整数
- key: 键名，字符串类型
- value: 键值，字符串类型
*/
type TestModel struct {
	ID    int    `gorm:"column:id" json:"id" form:"id" db:"id"`             // 主键ID，支持GORM、JSON、表单和数据库标签
	Key   string `gorm:"column:key" json:"key" form:"key" db:"key"`         // 键名字段，支持多种序列化格式
	Value string `gorm:"column:value" json:"value" form:"value" db:"value"` // 键值字段，支持多种序列化格式
}

// TableName 方法 - 指定数据库表名
// ElasticView插件SDK要求模型实现此方法来指定对应的数据库表名
// 返回值: 数据库表名字符串
func (this *TestModel) TableName() string {
	return "test" // 返回表名"test"
}

// Insert 方法 - 插入新记录到数据库
// 使用ElasticView插件SDK的SQL构建器生成INSERT语句并执行
// 返回值: 错误信息，nil表示操作成功
func (this *TestModel) Insert() (err error) {
	// 使用SQL构建器创建INSERT语句
	sql, args := sql_builder.SqlBuilder.
		Insert(this.TableName()).      // 指定插入的表名
		SetMap(map[string]interface{}{ // 设置插入的字段和值
			"`key`":   this.Key,   // 设置key字段值，使用反引号避免关键字冲突
			"`value`": this.Value, // 设置value字段值，使用反引号避免关键字冲突
		}).MustSql() // 生成SQL语句和参数
	// 执行SQL语句，使用ElasticView插件SDK的数据库执行接口
	_, err = ev_api.GetEvApi().StoreExec(context.Background(), sql, args...)
	if err != nil { // 检查执行是否出错
		return // 如果出错，返回错误信息
	}
	return // 操作成功，返回nil
}

// Clean 方法 - 清空表中所有数据
// 使用ElasticView插件SDK的SQL构建器生成DELETE语句并执行
// 返回值: 错误信息，nil表示操作成功
func (this *TestModel) Clean() (err error) {
	// 使用SQL构建器创建DELETE语句
	sql, args := sql_builder.SqlBuilder.
		Delete(this.TableName()). // 指定删除操作的表名
		MustSql()                 // 生成SQL语句和参数（无WHERE条件，删除所有记录）
	// 执行SQL语句，使用ElasticView插件SDK的数据库执行接口
	_, err = ev_api.GetEvApi().StoreExec(context.Background(), sql, args...)
	if err != nil { // 检查执行是否出错
		return // 如果出错，返回错误信息
	}
	return // 操作成功，返回nil
}

// List 方法 - 分页查询数据列表
// 使用ElasticView插件SDK的SQL构建器生成SELECT语句并执行分页查询
// 参数 page: 页码，从1开始
// 参数 limit: 每页记录数
// 返回值: 数据列表切片和错误信息
func (this *TestModel) List(page, limit int) (list []TestModel, err error) {
	// 使用SQL构建器创建SELECT查询语句
	builder := sql_builder.SqlBuilder.
		Select("*").                                // 选择所有字段
		From(this.TableName()).                     // 指定查询的表名
		OrderBy("id desc").                         // 按ID降序排列
		Limit(uint64(limit)).                       // 设置查询记录数限制
		Offset(sql_builder.CreatePage(page, limit)) // 设置分页偏移量

	// 生成最终的SQL语句和参数
	SQL, args, err := builder.ToSql()

	if err != nil { // 检查SQL生成是否出错
		return nil, err // 如果出错，返回错误信息
	}
	// 执行查询，使用ElasticView插件SDK的数据库查询接口
	err = ev_api.GetEvApi().StoreSelect(context.Background(), &list, SQL, args...)
	// 检查查询结果，如果是"无记录"错误则忽略，其他错误需要返回
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err // 如果是其他错误，返回错误信息
	}

	return list, nil // 返回查询结果列表
}

// Count 方法 - 统计表中记录总数
// 使用ElasticView插件SDK的SQL构建器生成COUNT语句并执行统计查询
// 返回值: 记录总数和错误信息
func (this *TestModel) Count() (count int, err error) {
	// 使用SQL构建器创建COUNT查询语句
	builder := sql_builder.SqlBuilder.
		Select("count(*) c").  // 选择记录总数，别名为c
		From(this.TableName()) // 指定查询的表名

	// 生成最终的SQL语句和参数
	sql, args, err := builder.ToSql()
	if err != nil { // 检查SQL生成是否出错
		return // 如果出错，返回错误信息
	}

	// 定义接收COUNT结果的结构体
	// 需要设置结构体来接收返回值
	type Cnt struct {
		C int `json:"c"` // COUNT结果字段，对应SQL中的别名c
	}
	var cnt Cnt // 创建结构体实例

	// 执行查询，使用ElasticView插件SDK的数据库单条记录查询接口
	err = ev_api.GetEvApi().StoreFirst(context.Background(), &cnt, sql, args...)

	if err != nil { // 检查查询是否出错
		return // 如果出错，返回错误信息
	}

	count = cnt.C // 提取COUNT结果

	return // 返回统计结果
}
