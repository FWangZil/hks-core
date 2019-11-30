package safeservice

import (
	"github.com/gin-gonic/gin"
)

// initRouter 初始化全部路由
func initRouter(router *gin.Engine) {
	// 总的文档首页在docs
	router.Static("/docs", "./docs")
	initNoAuthRouter(router)
}

func initNoAuthRouter(r *gin.Engine) {
	//apiRouter := r.Group("/api")
}
