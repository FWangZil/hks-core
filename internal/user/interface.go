package user

import (
	"hks/hks-core/pkg"

	"github.com/shopspring/decimal"
)

type repoI interface {
	// GetUserByQuery 获取用户详情
	GetUserByQuery(query Query) (*pkg.User, error)
	// ListUser 获取用户详情
	ListUser(query Query) ([]pkg.User, uint, error)
	// UserRegister 用户注册方法
	UserRegister(user *pkg.User) (*pkg.User, error)

	// About Locations
	// GetLocation 获取用户实时地址
	GetLocation()
	// ListLocations 列出用户实时地址
	ListLocations()

	GetUserStatus(longitude, latitude decimal.Decimal) (bool, error)
}
