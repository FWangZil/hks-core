package safeservice

import (
	"hks/hks-core/internal/conf"
	"hks/hks-core/internal/event"
	"hks/hks-core/internal/sql"
	"hks/hks-core/internal/user"
	"log"

	"github.com/gin-gonic/gin"
)

// Init 初始化
func Init() {

	// 初始化配置模块
	conf.Init() // 配置项初始化
	sql.Init()  // 数据库初始化

	user.Init()  // 用户模块初始化
	event.Init() // 事件模块初始化

	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	initRouter(r)
	log.Fatal(r.Run(":8080"))
}
