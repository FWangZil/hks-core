package shared

import (
	"github.com/meikeland/errkit"
)

// 常规错误
var (
	ErrUnLogin = errkit.New("未登录")
)
