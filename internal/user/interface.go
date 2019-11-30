package user

import "hks/hks-core/pkg"

type repoI interface {
	// GetUserByID 获取用户详情
	GetUserByID(userID uint) (*pkg.User, error)
	// ListUser 获取用户详情
	ListUser(query Query) ([]pkg.User, uint, error)
	// UserRegister 用户注册方法
	UserRegister(user *pkg.User) (*pkg.User, error)
	// GetLocation 获取用户实时地址
	GetLocation()
	// ListLocations 列出用户实时地址
	ListLocations()
}