package event

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Query 查询条件
type Query struct {
	Limit     uint
	Offset    uint
	EventID   uint // 事件自增ID
	UserID    uint
	Status    uint
	Level     uint
	StartTime time.Time
	EndTime   time.Time
	FixerID   uint
	HelperID  uint
}

// Where 条件
func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.EventID > 0 {
			db = db.Where("id = ?", c.EventID)
		}
		if c.UserID > 0 {
			db = db.Where("user_id = ?", c.UserID)
		}
		if c.Status > 0 {
			db = db.Where("status = ?", c.Status)
		}
		if c.FixerID > 0 {
			db = db.Where("fixer_id = ?", c.FixerID)
		}
		if c.HelperID > 0 {
			db = db.Where("helper_id = ?", c.HelperID)
		}
		if !c.StartTime.IsZero() && !c.EndTime.IsZero() {
			db = db.Where("time BETWEEN ? AND ?", c.StartTime, c.EndTime)
		}
		return db
	}
}
