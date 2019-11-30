package pkg

import (
	"errors"

	"github.com/FWangZil/errkit"
)

// VersionParam 创建新版本信息
type VersionParam struct {
	VersionID    uint   `json:"versionID"`    // 版本信息ID
	AppName      string `json:"appName"`      // app 名称
	Platform     string `json:"platform"`     // 平台（android or ios）
	VersionCode  uint   `json:"versionCode"`  // 版本号
	VersionLabel string `json:"versionLabel"` // 版本名称
	ChangeLog    string `json:"changeLog"`    // 更新内容
	IsForce      uint   `json:"isForce"`      // 是否强制更新
	URL          string `json:"url"`          // 下载地址
}

// IsValid 创建版本信息参数验证
func (param VersionParam) IsValid() error {
	if len(param.AppName) <= 0 {
		return errkit.New("App 名称不能为空")
	}

	if len(param.Platform) <= 0 {
		return errkit.New("平台信息 不能为空")
	}
	if param.VersionCode == 0 {
		return errkit.New("版本号 不能为空")
	}
	if len(param.VersionLabel) <= 0 {
		return errkit.New("版本名称 不能为空")
	}
	if len(param.ChangeLog) <= 0 {
		return errkit.New("更新内容 不能为空")
	}
	switch param.IsForce {
	case 1:
	case 2:
	default:
		return errors.New("无效的强制更新状态")
	}
	if len(param.URL) <= 0 {
		return errkit.New("下载地址 不能为空")
	}

	return nil
}
