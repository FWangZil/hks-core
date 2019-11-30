package opauth

import (
	"hks/hks-core/pkg"
	"hks/hks-core/util"

	"github.com/jinzhu/gorm"
	"github.com/meikeland/errkit"
)

type sqlRepo struct {
	db *gorm.DB
}

// Login 登录
func (repo sqlRepo) Login(account, password string) (*pkg.User, error) {
	query := query{
		Account: account,
	}

	user := &pkg.User{}
	if err := repo.db.Scopes(query.where()).First(user).Error; err != nil {
		return nil, errkit.Wrap(err, "获取管理员失败")
	}

	if ok := util.PasswordEncryCompare(user.Password, password); !ok {
		return nil, errkit.New("账号或密码错误")
	}
	return user, nil
}
