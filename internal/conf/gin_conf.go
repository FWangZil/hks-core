package conf

import "github.com/spf13/viper"

// Gin gin配置
type Gin struct {
	Mode string
}

// GetGin 获取gin配置
func GetGin() *Gin {
	return &Gin{
		Mode: viper.GetString("gin.mode"),
	}
}
