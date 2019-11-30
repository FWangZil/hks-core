package util

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"

	errkit "github.com/meikeland/errkit"
)

// ParseUint 转换字符串为uint
func ParseUint(key string) (uint, error) {
	if key == "" {
		return 0, nil
	}
	id, err := strconv.ParseUint(key, 10, strconv.IntSize)
	return uint(id), err
}

// TrimSpace 去空格
func TrimSpace(a string) string {
	a = strings.TrimSpace(a)
	a = strings.Replace(a, " ", "", -1)
	return a
}

// BubbleSort 冒泡排序
func BubbleSort(arr []uint) (res []uint) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		res = append(res, arr[i])
	}
	return
}

// Intersect 交集计算
func Intersect(arr1 []uint, arr2 []uint) []uint {
	if arr1 == nil || arr2 == nil {
		return []uint{}
	}
	BubbleSort(arr1)
	BubbleSort(arr2)
	var x, y = 0, 0
	var result []uint
	for {
		if x < len(arr1) && y < len(arr2) {
			if arr1[x] == arr2[y] {
				result = append(result, arr1[x])
				x++
				y++
			} else if arr1[x] > arr2[y] {
				y++
			} else {
				x++
			}
		} else {
			break
		}
	}
	return result
}

// Difference 差集计算
func Difference(arr1 []uint, arr2 []uint) []uint {
	if arr1 == nil && arr2 == nil {
		return []uint{}
	}
	if len(arr2) > len(arr1) {
		arr1, arr2 = arr2, arr1
	}
	var result []uint
	var isEqual bool
	for x := 0; x < len(arr1); x++ {
		for y := 0; y < len(arr2); y++ {
			isEqual = false
			if arr1[x] == arr2[y] {
				isEqual = true
				break
			}
		}
		if isEqual {
			continue
		}
		result = append(result, arr1[x])
	}
	return result
}

// GetHash 生成hash值
func GetHash() string {
	// 获取当前时间的时间戳
	t := time.Now().Unix()

	// 生成一个MD5的哈希
	h := md5.New()

	// 将时间戳转换为byte，并写入哈希
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(t))
	h.Write([]byte(b))

	// 将字节流转化为16进制的字符串
	return hex.EncodeToString(h.Sum(nil))
}

// 转换为我们规定的周天
func GetweekIndex(index uint) uint {
	var weekDayIndex uint
	switch index {
	case 0:
		weekDayIndex = 6
	default:
		weekDayIndex = index - 1
	}
	return weekDayIndex
}

// ValiateMobileNumber 校验手机号码是否有效
func ValiateMobileNumber(mobileNumber string) error {
	match, _ := regexp.MatchString("^1[23456789][0-9]{9}$", mobileNumber)
	if match {
		return nil
	}
	return errkit.Wrap(errkit.New("无效的手机号码"), "无效的手机号码")
}

// PasswordEncry 密码加密
func PasswordEncry(password string) (string, error) {
	password = TrimSpace(password)
	match, _ := regexp.MatchString("^[a-zA-Z0-9]{6,18}$", password)
	if !match {
		return "", errkit.Wrap(errkit.New("无效的密码"), "无效的密码，请使用6-18位字母和数字")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(TrimSpace(password)), bcrypt.DefaultCost)
	if err != nil {
		return "", errkit.Wrap(errkit.New("无法加密你设定的密码"), "无法加密你设定的密码")
	}
	return string(hashedPassword), nil
}

// PasswordEncryCompare 密码加密对比
func PasswordEncryCompare(password, compare string) bool {
	if err := ValiatePassword(compare); err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(TrimSpace(compare))); err != nil {
		return false
	}
	return true
}

// ValiatePassword 校验密码是否有效
func ValiatePassword(password string) error {
	password = TrimSpace(password)
	match, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", password)
	if match {
		return nil
	}
	return errors.Wrap(errors.New("无效的密码"), "无效的密码，请使用6-16位字母和数字")
}

// CheckIDCard 验证身份证号是否合法
func CheckIDCard(idCard string) error {
	res, _ := regexp.Match("^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$", []byte(idCard))
	if res {
		return nil
	}
	return errkit.Wrap(errkit.New("无效的身份证号"), "无效的身份证号")
}
