package pkg

import (
	"time"

	"github.com/shopspring/decimal"
)

// Event 事件
type Event struct {
	GormModel
	Time      time.Time       `json:"time" gorm:"type:date" form:"time"`              // 事件发生的事件
	Level     uint            `json:"level" gorm:"size:2" form:"level"`               // 事件等级
	Address   string          `json:"address" gorm:"size:200" form:"address"`         // 事情发生的地址
	Longitude decimal.Decimal `json:"longitude" gorm:"type:decimal" form:"longitude"` // 经度
	Latitude  decimal.Decimal `json:"latitude" gorm:"type:decimal" form:"latitude"`   // 维度
	Status    uint            `json:"status" gorm:"size:200" form:"status"`           // 事件发展至今的状态
	File      string          `json:"file" gorm:"type:text" form:"file"`              // 该事件的存档文件
	UserID    uint            `json:"userID" gorm:"size:11" form:"userID"`            // 报警人ID
	FixerID   uint            `json:"fixerID" gorm:"size:11" form:"fixerID"`          // 处理人ID 警察
	HelperID  uint            `json:"helperID" gorm:"size:11" form:"helperID"`        // 帮助人ID 保安 热心群众
}

// WarningArea 危险区域
type WarningArea struct {
	GormModel
	Longitude1 decimal.Decimal `json:"longitude1" gorm:"type:decimal"` // 经度1
	Latitude1  decimal.Decimal `json:"latitude1" gorm:"type:decimal"`  // 维度1
	Longitude2 decimal.Decimal `json:"longitude2" gorm:"type:decimal"` // 经度2
	Latitude2  decimal.Decimal `json:"latitude2" gorm:"type:decimal"`  // 维度2
	Level      uint            `json:"level" gorm:"size:2"`            // 事件等级
}
