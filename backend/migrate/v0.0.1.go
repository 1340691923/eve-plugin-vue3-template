package migrate

import (
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

func V0_0_1() *build.Migration {
	return &build.Migration{
		ID: "0.0.1",
		SqliteMigrateSqls: []*build.ExecSql{
			{
				//务必明确好
				Sql: `create table test
(
    id      INTEGER   not null primary key,
    key     TEXT    default '',
    value   TEXT 	default ''
);
`,
			},
		},
		MysqlMigrateSqls: []*build.ExecSql{
			{
				Sql: "CREATE TABLE test " +
					"(    id      int(11) NOT NULL AUTO_INCREMENT," +
					"   `key`  varchar(225)   DEFAULT ''," +
					"   `value`    varchar(255)  DEFAULT ''," +
					"    PRIMARY KEY (id) USING BTREE" +
					") ENGINE = InnoDB ;",
			},
		},
	}
}
