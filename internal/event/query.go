package event

import (
	"github.com/jinzhu/gorm"
)

// Query 查询条件
type Query struct {
	Limit   uint
	Offset  uint
	EventID uint // 事件自增ID
}

// Where 条件
func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.EventID > 0 {
			db = db.Where("id = ?", c.EventID)
		}
		return db
	}
}
