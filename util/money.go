package util

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// YuanToFenInt 元 to 分 int
func YuanToFenInt(a decimal.Decimal) int {
	f := 100.0
	return SystemDecimalToInt(a.Mul(decimal.NewFromFloat(f)))
}

// DecimalScoreToInt 积分decimal 转整形
func DecimalScoreToInt(a decimal.Decimal) int {
	// var f float64
	// f = 100.0
	return DecimalToInt(a, -1)
}

// SystemDecimalToInt deciaml to int
func SystemDecimalToInt(a decimal.Decimal) int {
	return DecimalToInt(a, 1)
}

// DecimalToInt deciaml to int,s:-1 floor;s:0 round;s:1 ceil
func DecimalToInt(a decimal.Decimal, s int) int {
	var temp string
	switch {
	case s == 0:
		temp = fmt.Sprintf("%s", a.Round(0))
	case s >= 1:
		temp = fmt.Sprintf("%s", a.Ceil())
	case s <= -1:
		temp = fmt.Sprintf("%s", a.Floor())
	}
	b, _ := strconv.Atoi(temp)
	return b
}
