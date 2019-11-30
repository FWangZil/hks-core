package safeservice

import (
	"hks/hks-core/internal/conf"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// initRouter 初始化全部路由
func initRouter(router *gin.Engine) {
	// 总的文档首页在docs

	router.Use(Cors())

	router.Static("/api/docs", "./docs")

	// 运营平台以session作为登录凭证
	router.Use(sessions.Sessions("safe", conf.GetSessionStore()))

	router.POST("/api/login", login)
	router.POST("/api/logout", logoutUser)
	router.GET("/api/user/register", registerUser)
	// 身份拦截
	//router.Use(auth)

	router.GET("/api/currentUser", currentUser) // 获取当前用户

	// initNeedAuthRouter 不需要登录的接口
	initNeedAuthRouter(router)

}

// 不需要登录的接口
func initNeedAuthRouter(r *gin.Engine) {
	user := r.Group("/api/user")
	userGroup := user.Group("/")
	{
		userGroup.GET("/get", getUserByQuery)
		userGroup.GET("/get/relationship", getRelationship)
		userGroup.GET("/list", listUser)
		userGroup.POST("/add/relationship", addRelationship)
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
