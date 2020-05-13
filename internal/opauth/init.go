package opauth

import "github.com/FWangZil/hks-core/internal/sql"

// ACL 权限控制器
var ACL acl

// Init 模块初始化
func Init() {
	ACL = sqlRepo{
		db: sql.Db,
	}
}
