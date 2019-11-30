package safeservice

import (
	"github.com/gin-gonic/gin"
)

// initRouter 初始化全部路由
func initRouter(router *gin.Engine) {
	// 总的文档首页在docs
	router.Static("/api/docs", "./docs")

	// 运营平台以session作为登录凭证
	//router.Use(sessions.Sessions("safe", conf.GetSessionStore()))

	router.POST("/login", login)
	router.POST("/logout", logoutUser)
	router.GET("/currentUser", currentUser) // 获取当前用户

	// 身份拦截
	router.Use(auth)

	// initNoAuthRouter 不需要登录的接口
	initNoAuthRouter(router)

}

// 不需要登录的接口
func initNoAuthRouter(r *gin.Engine) {
	// 统一登录地址
	// 统一注销地址

	user := r.Group("/api/user")
	//user.POST("/login")
	userGroup := user.Group("/")
	{
		userGroup.GET("/get", getUserByQuery)
		userGroup.GET("/get/relationship", getRelationship)
		userGroup.GET("/list", listUser)
		userGroup.POST("/register", registerUser)
		userGroup.PUT("/add/relationship", addRelationship)
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
		eventGroup.PUT("/update", updateEvent)
	}
}
