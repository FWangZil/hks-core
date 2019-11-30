package safeservice

import (
	"github.com/gin-gonic/gin"
)

// initRouter 初始化全部路由
func initRouter(router *gin.Engine) {
	// 总的文档首页在docs
	router.Static("/docs", "./docs")
	// initNoAuthRouter 不需要登录的接口
	initNoAuthRouter(router)
}

func initNoAuthRouter(r *gin.Engine) {
	user := r.Group("/api/user")
	userGroup := user.Group("/")
	{
		// 获取信息
		userGroup.GET("/get",getUserByID)
		userGroup.GET("/list",listUser)
		userGroup.GET("/register",registerUser)
	}
}
