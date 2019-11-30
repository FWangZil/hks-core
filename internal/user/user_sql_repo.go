package user

import (
	"fmt"
	"hks/hks-core/pkg"

	"github.com/FWangZil/errkit"

	"github.com/jinzhu/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

// GetUserByQuery 获取用户详情
func (repo sqlRepo) GetUserByQuery(query Query) (*pkg.User, error) {
	user := &pkg.User{}
	if err := repo.db.Model(user).Scopes(query.where()).First(user).Error; err != nil {
		return nil, fmt.Errorf("获取用户信息发生错误：%w", err)
	}
	return user, nil
}

// ListUser 获取用户详情
func (repo sqlRepo) ListUser(query Query) ([]pkg.User, uint, error) {
	count, err := repo.count(query)
	if err != nil {
		return nil, 0, err
	}
	var users []pkg.User
	if err := repo.db.Model(&pkg.User{}).Scopes(query.where()).Offset(query.Offset).Limit(query.Limit).Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("获取用户信息发生错误：%w", err)
	}
	return users, count, nil
}

// UserRegister 用户注册方法
func (repo sqlRepo) UserRegister(user *pkg.User) (*pkg.User, error) {
	if err := repo.db.Model(&pkg.User{}).Create(user).Error; err != nil {
		return nil, fmt.Errorf("注册用户信息发生错误：%w", err)
	}
	return user, nil
}

// AddRelationship 增加用户亲友关系
func (repo sqlRepo) AddRelationship(relation *pkg.Relative) (*pkg.Relative, error) {
	if err := repo.db.Model(&pkg.Relative{}).Create(relation).Error; err != nil {
		return nil, fmt.Errorf("注册用户信息发生错误：%w", err)
	}
	return relation, nil
}

// AddRelationship 获取用户亲友关系列表
func (repo sqlRepo) GetUserRelationship(userID uint) ([]pkg.Relative, error) {
	var relativeArr []pkg.Relative
	if err := repo.db.Model(&pkg.Relative{}).Where("user_id = ?", userID).Find(&relativeArr).Error; err != nil {
		return nil, fmt.Errorf("注册用户信息发生错误：%w", err)
	}
	return relativeArr, nil
}

// count 获取用户记录数量
func (repo sqlRepo) count(query Query) (uint, error) {
	var count uint
	if err := repo.db.Model(&pkg.User{}).Scopes(query.where()).Count(&count).Error; err != nil {
		return 0, errkit.Wrap(err, "获取用户量失败")
	}
	return count, nil
}
