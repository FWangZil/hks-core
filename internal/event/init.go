package event

import (
	"hks/hks-core/internal/sql"
)

var (
	Repo repoI
)

func Init() {
	Repo = sqlRepo{
		db: sql.Db,
	}
}
