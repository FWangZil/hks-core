package sql

import (
	"hks/hks-core/internal/conf"
	"hks/hks-core/pkg"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 隐式使用mysql
	"github.com/meikeland/errkit"
)

var (
	// Db gorm的数据库连接
	Db *gorm.DB
)

// Init 初始化函数
func Init() {
	database := conf.GetDatabase()

	var err error
	Db, err = gorm.Open(database.Dialect, database.Dsn)
	if err != nil {
		panic(errkit.Wrapf(err, "failed to connect database, datebase config is %v", database))
	}
	if database.Debug {
		Db = Db.Debug()
	}
	Db.AutoMigrate(
		&pkg.User{},
		&pkg.UserLocation{},
		&pkg.Event{},
		&pkg.WarningArea{},
	)
	log.Print("All table AutoMigrate finish.")
}
