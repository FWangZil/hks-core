package pkg

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	GenderMan    = 1
	GenderFemale = 2
)

const (
	TypeUser      = 1
	TypePolice    = 2
	TypeDetective = 3
)

const (
	StatusWaiting  = 1
	StatusDelivery = 2
	StatusPolice   = 3
	StatusEnd      = 4
)

// User 用户
type User struct {
	GormModel
	Name         string          `json:"name" gorm:"size:100"`          // 用户姓名
	Mobile       string          `json:"mobile" gorm:"size:11"`         // 手机号码
	Password     string          `json:"password" gorm:"size:200"`      // 用户登录密码
	IDCard       string          `json:"idCard" gorm:"size:200"`        // 身份证号码 警员标志码
	Gender       uint            `json:"gender" gorm:"size:2"`          // 性别
	Type         uint            `json:"type" gorm:"size:2"`            // 用户类型
	Address      string          `json:"address" gorm:"size:200"`       // 警察的话是警局 保安时负责街区 用户是常用活动范围
	Longitude    decimal.Decimal `json:"longitude" gorm:"type:decimal"` // 经度
	Latitude     decimal.Decimal `json:"latitude" gorm:"type:decimal"`  // 维度
	Photo        string          `json:"photo" gorm:"type:text"`        // 照片
	UserLocation *UserLocation   `json:"userLocation,omitempty"`        // 用户的实时位置
}

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

// UserLocation 用户的实时位置点
type UserLocation struct {
	GormModel
	UserID    uint            `json:"userID" gorm:"size:11"`
	Longitude decimal.Decimal `json:"longitude" gorm:"type:decimal"` // 经度
	Latitude  decimal.Decimal `json:"latitude" gorm:"type:decimal"`  // 维度
}
