package pkg

import (
	"regexp"

	"github.com/FWangZil/errkit"
)

// ValiateMobileNumber 校验手机号码是否有效
func ValiateMobileNumber(mobileNumber string) error {
	match, _ := regexp.MatchString("^1[23456789][0-9]{9}$", mobileNumber)
	if match {
		return nil
	}
	return errkit.Wrap(errkit.New("无效的手机号码"), "无效的手机号码")
}

// CheckEmail 检查邮箱是否合法
func CheckEmail(email string) error {
	match, _ := regexp.MatchString("^[a-z0-9A-Z]+[- | a-z0-9A-Z . _]+@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-z]{2,}$", email)
	if match {
		return nil
	}
	return errkit.Wrap(errkit.New("无效的邮箱账号"), "无效的邮箱账号")
}

// CheckIDCard 验证身份证号是否合法
func CheckIDCard(idCard string) error {
	res, _ := regexp.Match("^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$", []byte(idCard))
	if res {
		return nil
	}
	return errkit.Wrap(errkit.New("无效的身份证号"), "无效的身份证号")
}
