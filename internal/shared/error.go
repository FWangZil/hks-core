package shared

import (
	"github.com/FWangZil/errkit"
)

// 常规错误
var (
	ErrUnLogin = errkit.New("未登录")
)
