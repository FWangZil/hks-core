package safeservice

import (
	"hks/hks-core/internal/conf"
	"hks/hks-core/internal/sql"
	"log"

	"github.com/gin-gonic/gin"
)

// Init 初始化
func Init() {

	// 初始化配置模块
	conf.Init() // 配置项初始化
	sql.Init()  // 数据库初始化

	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	initRouter(r)
	log.Fatal(r.Run(":8080"))
}
