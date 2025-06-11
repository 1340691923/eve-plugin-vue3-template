// ElasticView插件数据库迁移包 - 管理插件数据库表结构的版本化更新
package migrate

import (
	"github.com/1340691923/eve-plugin-sdk-go/build" // 导入ElasticView插件SDK构建工具包
)

// V0_0_1 函数 - 创建版本0.0.1的数据库迁移配置
// ElasticView插件支持多种数据库类型，需要为每种数据库提供对应的SQL语句
// 返回值: 包含迁移SQL语句的Migration实例
func V0_0_1() *build.Migration {
	return &build.Migration{
		ID: "0.0.1", // 迁移版本号，用于标识和管理数据库表结构版本
		// SQLite数据库的迁移SQL语句列表
		// ElasticView插件默认使用SQLite作为轻量级存储方案
		SqliteMigrateSqls: []*build.ExecSql{
			{
				// SQLite建表语句 - 创建test表
				// 注意：SQLite语法相对简单，主键使用INTEGER类型
				Sql: `create table test
(
    id      INTEGER   not null primary key,
    key     TEXT    default '',
    value   TEXT 	default ''
);
`,
			},
		},
		// MySQL数据库的迁移SQL语句列表
		// 当ElasticView插件配置使用MySQL时会执行这些语句
		MysqlMigrateSqls: []*build.ExecSql{
			{
				// MySQL建表语句 - 创建test表
				// 注意：MySQL语法更加完善，支持AUTO_INCREMENT和存储引擎配置
				Sql: "CREATE TABLE test " +
					"(    id      int(11) NOT NULL AUTO_INCREMENT," + // 主键ID，自增整数
					"   `key`  varchar(225)   DEFAULT ''," + // 键名字段，使用反引号避免关键字冲突
					"   `value`    varchar(255)  DEFAULT ''," + // 键值字段，使用反引号避免关键字冲突
					"    PRIMARY KEY (id) USING BTREE" + // 设置主键和索引类型
					") ENGINE = InnoDB ;", // 使用InnoDB存储引擎
			},
		},
	}
}
