package safeservice

import (
	"github.com/FWangZil/hks-core/internal/conf"
	"github.com/FWangZil/hks-core/internal/event"
	"github.com/FWangZil/hks-core/internal/opauth"
	"github.com/FWangZil/hks-core/internal/sql"
	"github.com/FWangZil/hks-core/internal/user"
	"log"

	"github.com/gin-gonic/gin"
)

// Init 初始化
func Init() {

	// 初始化配置模块
	conf.Init() // 配置项初始化
	sql.Init()  // 数据库初始化

	user.Init()   // 用户模块初始化
	event.Init()  // 事件模块初始化
	opauth.Init() // 通知模块初始化

	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	initRouter(r)
	log.Fatal(r.Run(":8080"))
}
