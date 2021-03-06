package safeservice

import (
	"fmt"
	"github.com/FWangZil/hks-core/internal/event"
	"github.com/FWangZil/hks-core/pkg"
	"github.com/FWangZil/hks-core/util"
	"log"
	"time"

	"github.com/FWangZil/errkit"
	"github.com/gin-gonic/gin"
)

// getEventByID 通过ID获取事件信息
func getEventByID(c *gin.Context) {
	eventID, err := util.ParseUint(c.Query("eventID"))
	if err != nil {
		fail(c, err)
		return
	}
	eventInfo, err := event.Repo.GetEventByID(eventID)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": eventInfo,
	})
}

// getNewestEventByUserID 获取某用户最新的报警事件
func getNewestEventByUserID(c *gin.Context) {
	userID, err := util.ParseUint(c.Query("userID"))
	if err != nil {
		fail(c, err)
		return
	}
	eventInfo, err := event.Repo.GetNewestEventByUserID(userID)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": eventInfo,
	})
}

// listEvent 通过ID获取事件信息
func listEvent(c *gin.Context) {
	param := struct {
		EventID   uint   `json:"eventID" form:"eventID"`     // 事件自增ID
		UserID    uint   `json:"userID" form:"userID"`       // 事件自增ID
		Gender    uint   `json:"gender" form:"gender"`       // 性别
		Type      uint   `json:"type" form:"type"`           // 事件类型
		StartTime string `json:"startTime" form:"startTime"` // 筛选开始事件
		EndTime   string `json:"endTime" form:"endTime"`     // 筛选结束事件
		FixerID   uint   `json:"fixerID" form:"fixerID"`
		HelperID  uint   `json:"helperID" form:"helperID"`
	}{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", param.StartTime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", param.EndTime, time.Local)
	pagination := paginationByGet(c)
	query := event.Query{
		EventID:   param.EventID,
		Limit:     pagination.PageSize,
		Offset:    pagination.Offset,
		UserID:    param.UserID,
		StartTime: startTime,
		EndTime:   endTime,
		FixerID:   param.FixerID,
		HelperID:  param.HelperID,
	}
	events, count, err := event.Repo.ListEvent(query)
	if err != nil {
		fail(c, err)
		return
	}
	pagination.Total = count
	okArr(c, resp{
		"data":       events,
		"pagination": pagination,
	})
}

// registerEvent 事件注册接口
func registerEvent(c *gin.Context) {
	param := pkg.Event{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	param.Time = time.Now()
	eventInfo, err := event.Repo.EventRegister(&param)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": eventInfo,
	})
}

// updateEvent 事件更新接口
func updateEvent(c *gin.Context) {
	param := pkg.Event{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	param.Time = time.Now()
	eventInfo, err := event.Repo.UpdateEvent(&param)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": eventInfo,
	})
}

func deleteAllEvents(c *gin.Context) {
	err := event.Repo.DeleteAllEvents()
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{})
}
