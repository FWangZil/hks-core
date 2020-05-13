package user

import (
	"github.com/FWangZil/hks-core/internal/sql"
)

var (
	Repo repoI
)

func Init() {
	Repo = sqlRepo{
		db: sql.Db,
	}
}
