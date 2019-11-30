package user

import (
	"github.com/jinzhu/gorm"
)

// Query 查询条件
type Query struct {
	Limit  uint
	Offset uint
	UserID uint // 用户自增ID
	Mobile string
	Gender uint // 性别
	Type   uint // 用户类型
}

// Where 条件
func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.UserID > 0 {
			db = db.Where("id = ?", c.UserID)
		}
		if c.Gender > 0 {
			db = db.Where("gender = ?", c.Gender)
		}
		if c.Type > 0 {
			db = db.Where("type = ?", c.Type)
		}
		if len(c.Mobile) > 0 {
			db = db.Where("mobile = ?", c.Mobile)
		}
		return db
	}
}
