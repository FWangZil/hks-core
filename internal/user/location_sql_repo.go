package user

import (
	"hks/hks-core/pkg"
	"math"
	"time"

	"github.com/shopspring/decimal"
)

// GetLocation 获取用户实时地址
func (repo sqlRepo) GetLocation(userID uint) (*pkg.UserLocation, error) {
	ul := pkg.UserLocation{}
	if err := repo.db.Model(&pkg.UserLocation{}).Where("user_id = ? and type = ?", userID, pkg.NowLocations).First(&ul).Error; err != nil {
		return nil, err
	}
	return &ul, nil
}

// SetUserLocation 更新用户实时地址
func (repo sqlRepo) SetUserLocation(userID uint, longitude, latitude decimal.Decimal) (*pkg.UserLocation, error) {
	ul := pkg.UserLocation{}
	if err := repo.db.Model(&pkg.UserLocation{}).Where("user_id = ? and type = ?", userID, pkg.NowLocations).First(&ul).Error; err != nil {
		//return nil, err
	}
	if err := repo.db.Model(&pkg.UserLocation{}).Where("id = ?", ul.ID).Update("type", pkg.HistoryLocation).Error; err != nil {
		//return nil, err
	}
	cul := pkg.UserLocation{
		UserID:    userID,
		Time:      time.Now(),
		Longitude: longitude,
		Latitude:  latitude,
		Type:      pkg.NowLocations,
	}
	if err := repo.db.Model(&pkg.UserLocation{}).Create(&cul).Error; err != nil {
		return nil, err
	}
	return &cul, nil
}

// ListLocations 列出用户实时地址
func (repo sqlRepo) ListUserLocations(userID uint) ([]*pkg.UserLocation, error) {
	var uls []*pkg.UserLocation
	if err := repo.db.Model(&pkg.UserLocation{}).Where("user_id=?", userID).Find(&uls).Error; err != nil {
		return nil, err
	}
	return uls, nil
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
		distance := repo.getDistance(lat1, lat2, lng1, lng2)
		if distance > 10 {
			return true, nil
		}
	}
	return false, nil
}

//返回单位为：千米
func (repo sqlRepo) getDistance(lat1, lat2, lng1, lng2 float64) float64 {
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
