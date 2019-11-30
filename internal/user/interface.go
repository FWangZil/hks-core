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
	GetLocation(userID uint) (*pkg.UserLocation, error)
	// SetUserLocation 更新用户实时地址
	SetUserLocation(userID uint, longitude, latitude decimal.Decimal) (*pkg.UserLocation, error)
	// ListLocations 列出用户实时地址
	ListUserLocations(userID uint) ([]*pkg.UserLocation, error)
	// GetUserStatus 获取用户目前所在区域的危险情况
	GetUserStatus(longitude, latitude decimal.Decimal) (bool, error)
}
