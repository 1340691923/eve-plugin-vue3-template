package versions

import (
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

func init() {
	Register(V0_0_1)
}

func V0_0_1() *build.Migration {
	return &build.Migration{
		ID: "0.0.1",
		MigrateSqls: []*build.ExecSql{
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
	}
}
