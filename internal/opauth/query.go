package opauth

import "github.com/jinzhu/gorm"

// query 查询用户
type query struct {
	ID     uint
	Mobile string
}

// where 条件
func (c query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.ID > 0 {
			db = db.Where("id = ?", c.ID)
		}
		if len(c.Mobile) > 0 {
			db = db.Where("mobile = ?", c.Mobile)
		}
		return db
	}
}
