package pkg

import (
	"time"

	"github.com/shopspring/decimal"
)

// User 用户
type User struct {
	GormModel
	Name            string          `json:"name" gorm:"size:100" form:"name"`                 // 用户姓名
	Mobile          string          `json:"mobile" gorm:"size:11" form:"mobile"`              // 手机号码
	Password        string          `json:"password" gorm:"size:200" form:"password"`         // 用户登录密码
	IDCard          string          `json:"idCard" gorm:"size:200" form:"idCard"`             // 身份证号码 警员标志码
	Gender          uint            `json:"gender" gorm:"size:2" form:"gender"`               // 性别
	Age             uint            `json:"age" gorm:"size:10" form:"age"`                    // 用户年龄
	Type            uint            `json:"type" gorm:"size:2" form:"type"`                   // 用户类型
	Address         string          `json:"address" gorm:"size:200" form:"address"`           // 警察的话是警局 保安是负责街区 用户是常用活动范围
	Longitude       decimal.Decimal `json:"longitude" gorm:"type:decimal" form:"longitude"`   // 经度
	Latitude        decimal.Decimal `json:"latitude" gorm:"type:decimal" form:"latitude"`     // 维度
	Photo           string          `json:"photo" gorm:"type:text" form:"photo"`              // 照片
	UserLocationArr []UserLocation  `json:"userLocationArr,omitempty" form:"userLocationArr"` // 用户的实时位置
	RelativeArr     []Relative      `json:"relativeArr,omitempty" form:"relativeArr"`         // 亲友关系表
}

// UserLocation 用户的实时位置点
type UserLocation struct {
	GormModel
	UserID    uint            `json:"userID" gorm:"size:11" form:"userID"`            // 用户ID
	Time      time.Time       `json:"time" gorm:"type:date" form:"time"`              // 轨迹点记录时间
	Longitude decimal.Decimal `json:"longitude" gorm:"type:decimal" form:"longitude"` // 经度
	Latitude  decimal.Decimal `json:"latitude" gorm:"type:decimal" form:"latitude"`   // 维度
	Type      uint            `json:"type" gorm:"size:2" form:"type"`                 // 该位置嗲状态
}

// Relative 亲友
type Relative struct {
	GormModel
	UserID       uint   `json:"userID" gorm:"size:11" form:"userID"`               // 用户ID
	Mobile       string `json:"mobile" gorm:"size:11" form:"mobile"`               // 手机号码
	Relationship string `json:"relationship" gorm:"size:100"  form:"relationship"` // 关系
}
