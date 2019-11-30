package event

import (
	"fmt"
	"hks/hks-core/pkg"
	"time"

	"github.com/FWangZil/errkit"

	"github.com/jinzhu/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

// GetEventByID 获取事件详情
func (repo sqlRepo) GetEventByID(eventID uint) (*pkg.Event, error) {
	event := &pkg.Event{}
	if err := repo.db.Model(event).Where("id = ?", eventID).First(event).Error; err != nil {
		return nil, fmt.Errorf("获取事件信息发生错误：%w", err)
	}
	return event, nil
}

// ListEvent 获取事件详情
func (repo sqlRepo) ListEvent(query Query) ([]pkg.Event, uint, error) {
	count, err := repo.count(query)
	if err != nil {
		return nil, 0, err
	}
	var events []pkg.Event
	if err := repo.db.Model(&pkg.Event{}).Scopes(query.where()).Offset(query.Offset).Limit(query.Limit).Find(&events).Error; err != nil {
		return nil, 0, fmt.Errorf("获取事件信息发生错误：%w", err)
	}
	return events, count, nil
}

// EventRegister 事件注册方法
func (repo sqlRepo) EventRegister(event *pkg.Event) (*pkg.Event, error) {
	event.Status = pkg.StatusWaiting
	event.Time = time.Now()
	if err := repo.db.Model(&pkg.Event{}).Create(&event).Error; err != nil {
		return nil, fmt.Errorf("注册事件信息发生错误：%w", err)
	}
	return event, nil
}

// UpdateEvent 更新事件状态
func (repo sqlRepo) UpdateEvent(event *pkg.Event) (*pkg.Event, error) {
	if err := repo.db.Model(&pkg.Event{}).Update(&event).Error; err != nil {
		return nil, fmt.Errorf("注册事件信息发生错误：%w", err)
	}
	return event, nil
}

// count 获取事件记录数量
func (repo sqlRepo) count(query Query) (uint, error) {
	var count uint
	if err := repo.db.Model(&pkg.Event{}).Scopes(query.where()).Count(&count).Error; err != nil {
		return 0, errkit.Wrap(err, "获取事件量失败")
	}
	return count, nil
}
