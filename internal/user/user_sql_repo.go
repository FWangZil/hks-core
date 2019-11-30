package user

import (
	"fmt"
	"hks/hks-core/pkg"
	"math"

	"github.com/shopspring/decimal"

	"github.com/meikeland/errkit"

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
	if err := repo.db.Model(&pkg.User{}).Create(&user).Error; err != nil {
		return nil, fmt.Errorf("注册用户信息发生错误：%w", err)
	}
	return user, nil
}

// GetLocation 获取用户实时地址
func (repo sqlRepo) GetLocation(userID uint) (*pkg.UserLocation, error) {

}

// ListLocations 列出用户实时地址
func (repo sqlRepo) ListLocations() {

}

// GetUserStatus 获取用户目前所在区域的危险情况
func (repo sqlRepo) GetUserStatus(longitude, latitude decimal.Decimal) (bool, error) {
	var wa []pkg.WarningArea
	if err := repo.db.Model(&pkg.WarningArea{}).Find(&wa).Error; err != nil {
		return false, err
	}
	for _, v := range wa {
		longitude0 := v.Longitude1.Sub(v.Longitude2).Div(decimal.New(2, 0)).Abs()
		Latitude0 := v.Latitude1.Sub(v.Latitude2).Div(decimal.New(2, 0)).Abs()
		lat1, _ := latitude.Float64()
		lat2, _ := Latitude0.Float64()
		lng1, _ := longitude.Float64()
		lng2, _ := longitude0.Float64()
		distance := repo.GetDistance(lat1, lat2, lng1, lng2)
		if distance > 10 {
			return true, nil
		}
	}
	return false, nil
}

// count 获取用户记录数量
func (repo sqlRepo) count(query Query) (uint, error) {
	var count uint
	if err := repo.db.Model(&pkg.User{}).Scopes(query.where()).Count(&count).Error; err != nil {
		return 0, errkit.Wrap(err, "获取用户量失败")
	}
	return count, nil
}

//返回单位为：千米
func (repo sqlRepo) GetDistance(lat1, lat2, lng1, lng2 float64) float64 {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / 1000
}
