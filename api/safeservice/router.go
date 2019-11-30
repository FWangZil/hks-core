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

// 不需要登录的接口
func initNoAuthRouter(r *gin.Engine) {
	user := r.Group("/api/user")
	userGroup := user.Group("/")
	{
		userGroup.GET("/get", getUserByQuery)
		userGroup.GET("/list", listUser)
		userGroup.GET("/register", registerUser)
	}
	event := r.Group("/api/event")
	eventGroup := event.Group("/")
	{
		eventGroup.GET("/get", getEventByID)
		eventGroup.GET("/list", listEvent)
		eventGroup.GET("/register", registerEvent)
	}
}
