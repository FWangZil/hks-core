package event

import "github.com/FWangZil/hks-core/pkg"

type repoI interface {
	// GetEventByID 获取事件详情
	GetEventByID(eventID uint) (*pkg.Event, error)

	// GetNewestEventByUserID 获取事件详情
	GetNewestEventByUserID(userID uint) (*pkg.Event, error)

	// ListEvent 获取事件详情
	ListEvent(query Query) ([]pkg.Event, uint, error)
	// EventRegister 事件注册方法
	EventRegister(event *pkg.Event) (*pkg.Event, error)

	// UpdateEvent 更新事件状态
	UpdateEvent(event *pkg.Event) (*pkg.Event, error)

	DeleteAllEvents() error
}
