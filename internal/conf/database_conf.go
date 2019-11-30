package conf

import "github.com/spf13/viper"

// Database 数据库配置
type Database struct {
	Dialect string
	Dsn     string
	Debug   bool
}

// GetDatabase 获取数据库配置
func GetDatabase() *Database {
	return &Database{
		Dialect: viper.GetString("database.dialect"),
		Dsn:     viper.GetString("database.dsn"),
		Debug:   viper.GetBool("database.debug"),
	}
}
