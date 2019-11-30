package safeservice

import (
	"github.com/gin-gonic/gin"
)

// initRouter 初始化全部路由
func initRouter(router *gin.Engine) {
	// 总的文档首页在docs
	router.Static("/api/docs", "./docs")
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
		userGroup.POST("/register", registerUser)
	}
	location := user.Group("/location")
	{
		location.GET("/status", getUserStatus)
		location.GET("/get", getLocation)
		location.GET("/list", listUserLocations)
		location.POST("/set", setUserLocation)
	}
	event := r.Group("/api/event")
	eventGroup := event.Group("/")
	{
		eventGroup.GET("/get", getEventByID)
		eventGroup.GET("/list", listEvent)
		eventGroup.POST("/register", registerEvent)
	}
}
