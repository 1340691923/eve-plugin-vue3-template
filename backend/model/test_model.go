// 数据库实体层
package model

import (
	"context"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/sql_builder"
)

/*
	http://sql2struct.atotoa.com 根据表结构生成 go结构体

TestModel DSL历史记录
*/
type TestModel struct {
	ID    int    `gorm:"column:id" json:"id" form:"id" db:"id"`
	Key   string `gorm:"column:key" json:"key" form:"key" db:"key"`
	Value string `gorm:"column:value" json:"value" form:"value" db:"value"`
}

// 表名
func (this *TestModel) TableName() string {
	return "test"
}

// Insert
func (this *TestModel) Insert() (err error) {
	sql, args := sql_builder.SqlBuilder.
		Insert(this.TableName()).
		SetMap(map[string]interface{}{
			"`key`":   this.Key,
			"`value`": this.Value,
		}).MustSql()
	_, err = ev_api.GetEvApi().StoreExec(context.Background(), sql, args...)
	if err != nil {
		return
	}
	return
}

// Clean
func (this *TestModel) Clean() (err error) {
	sql, args := sql_builder.SqlBuilder.
		Delete(this.TableName()).
		MustSql()
	_, err = ev_api.GetEvApi().StoreExec(context.Background(), sql, args...)
	if err != nil {
		return
	}
	return
}

// List
func (this *TestModel) List(page, limit int) (list []TestModel, err error) {
	builder := sql_builder.SqlBuilder.
		Select("*").
		From(this.TableName()).
		OrderBy("id desc").
		Limit(uint64(limit)).Offset(sql_builder.CreatePage(page, limit))

	SQL, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}
	err = ev_api.GetEvApi().StoreSelect(context.Background(), &list, SQL, args...)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	return list, nil
}

// Count
func (this *TestModel) Count() (count int, err error) {
	builder := sql_builder.SqlBuilder.
		Select("count(*) c").
		From(this.TableName())

	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}

	//需要设置结构体接收返回值
	type Cnt struct {
		C int `json:"c"`
	}
	var cnt Cnt

	err = ev_api.GetEvApi().StoreFirst(context.Background(), &cnt, sql, args...)

	if err != nil {
		return
	}

	count = cnt.C

	return
}
