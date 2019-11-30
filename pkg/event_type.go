package pkg

import (
	"time"

	"github.com/shopspring/decimal"
)

// Event 事件
type Event struct {
	GormModel
	Time      time.Time       `json:"time" gorm:"type:date"`         // 事件发生的事件
	Level     uint            `json:"level" gorm:"size:2"`           // 事件等级
	Address   string          `json:"address" gorm:"size:200"`       // 事情发生的地址
	Longitude decimal.Decimal `json:"longitude" gorm:"type:decimal"` // 经度
	Latitude  decimal.Decimal `json:"latitude" gorm:"type:decimal"`  // 维度
	Status    string          `json:"status" gorm:"size:200"`        // 事件发展至今的状态
	File      string          `json:"file" gorm:"type:file"`         // 该事件的存档文件
}
