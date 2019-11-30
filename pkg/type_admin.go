package pkg

// Admin 管理员
type Admin struct {
	GormModel
	Nickname string `gorm:"size:20;" json:"nickname"` // 昵称
	Account  string `gorm:"size:20;" json:"account"`  // 帐号
	Password string `gorm:"size:100;" json:"-"`       // 密码
}
