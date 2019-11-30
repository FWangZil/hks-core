package safeservice

import (
	"fmt"
	"hks/hks-core/internal/user"
	"hks/hks-core/util"
	"log"

	"github.com/shopspring/decimal"

	"github.com/FWangZil/errkit"
	"github.com/gin-gonic/gin"
)

// getUserStatus 通过ID获取用户信息
func getUserStatus(c *gin.Context) {
	param := struct {
		UserID    uint            `json:"userID" form:"userID"`
		Longitude decimal.Decimal `json:"longitude" form:"longitude"` // 经度
		Latitude  decimal.Decimal `json:"latitude" form:"latitude"`   // 维度
		Mobile    string          `json:"mobile" form:"mobile"`
	}{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	userInfo, err := user.Repo.GetUserStatus(param.Longitude, param.Latitude)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"user": userInfo,
	})
}

// getLocation 获取用户当前实时地址
func getLocation(c *gin.Context) {
	userID, err := util.ParseUint(c.Query("userID"))
	if err != nil {
		fail(c, err)
		return
	}
	ul, err := user.Repo.GetLocation(userID)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"userLocation": ul,
	})
}

// listUserLocations 获取用户当前地址
func listUserLocations(c *gin.Context) {
	userID, err := util.ParseUint(c.Query("userID"))
	if err != nil {
		fail(c, err)
		return
	}
	uls, err := user.Repo.ListUserLocations(userID)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"userLocations": uls,
	})
}

// setUserLocation 更新用户实时地址
func setUserLocation(c *gin.Context) {
	param := struct {
		UserID    uint            `json:"userID"`
		Longitude decimal.Decimal `json:"longitude"` // 经度
		Latitude  decimal.Decimal `json:"latitude"`  // 维度
	}{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	ul, err := user.Repo.SetUserLocation(param.UserID, param.Longitude, param.Latitude)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"userLocation": ul,
	})
}
