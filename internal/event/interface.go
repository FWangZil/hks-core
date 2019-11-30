package event

import "hks/hks-core/pkg"

type repoI interface {
	// GetEventByID 获取事件详情
	GetEventByID(eventID uint) (*pkg.Event, error)
	// ListEvent 获取事件详情
	ListEvent(query Query) ([]pkg.Event, uint, error)
	// EventRegister 事件注册方法
	EventRegister(event *pkg.Event) (*pkg.Event, error)
}
