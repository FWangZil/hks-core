package opauth

import "hks/hks-core/pkg"

type acl interface {
	// 登录
	Login(account, password string) (*pkg.User, error)
}
