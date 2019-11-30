package opauth

import (
	"hks/hks-core/pkg"

	"github.com/FWangZil/errkit"
	"github.com/jinzhu/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

// Login 登录
func (repo sqlRepo) Login(account, password string) (*pkg.User, error) {
	query := query{
		Mobile: account,
	}
	user := &pkg.User{}
	if err := repo.db.Scopes(query.where()).First(user).Error; err != nil {
		return nil, errkit.Wrap(err, "获取用户信息失败")
	}
	if user.Password != password {
		return nil, errkit.New("密码错误")
	}
	return user, nil
}
