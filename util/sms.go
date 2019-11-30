package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"code.aliyun.com/yjqu/apigate-card/internal/conf"
	errkit "github.com/meikeland/errkit"
)

// YunPianTemplate 模板
type YunPianTemplate struct {
	ID      string
	Content string
}

// VerifyCode 生成六位数验证码
func VerifyCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	verifyCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return verifyCode
}

// exchangeConsumeURLValues 转换内容 2
func exchangeConsumeURLValues(store, amount, code string) url.Values {
	urlValues := url.Values{}
	// urlValues[fmt.Sprintf("#%s#", "store")] = []string{store}
	urlValues[fmt.Sprintf("#%s#", "amount")] = []string{amount}
	urlValues[fmt.Sprintf("#%s#", "code")] = []string{code}
	return urlValues
}

// exchangeRechargeURLValues 转换内容 1
func exchangeRechargeURLValues(amount, bonus, balance string) url.Values {
	urlValues := url.Values{}
	urlValues[fmt.Sprintf("#%s#", "amount")] = []string{amount}
	urlValues[fmt.Sprintf("#%s#", "bonus")] = []string{bonus}
	// urlValues[fmt.Sprintf("#%s#", "balance")] = []string{balance}
	return urlValues
}

// exchangeConsumeNotifyURLValues 转换内容 3
func exchangeConsumeNotifyURLValues(store, amount, balance string) url.Values {
	urlValues := url.Values{}
	urlValues[fmt.Sprintf("#%s#", "store")] = []string{store}
	urlValues[fmt.Sprintf("#%s#", "amount")] = []string{amount}
	// urlValues[fmt.Sprintf("#%s#", "balance")] = []string{balance}
	return urlValues
}

// exchangeTransferURLValues 转换内容 4
func exchangeTransferURLValues(amount string) url.Values {
	urlValues := url.Values{}
	urlValues[fmt.Sprintf("#%s#", "amount")] = []string{amount}
	return urlValues
}

// SendCodeYunPian 云片发送验证码
func SendCodeYunPian(mobile, store, amount, bonus, balance string, kind uint) (string, error) {
	var sms *conf.SMS
	var values url.Values
	verifyCode := VerifyCode()
	switch kind {
	// TODO：待重构为常量
	case 1:
		// 充值通知
		values = exchangeRechargeURLValues(amount, bonus, balance)
		sms = conf.GetRechargeSMS()
	case 2:
		// 消费验证
		log.Println("sms", verifyCode)
		log.Println("mobileNumber", mobile)
		values = exchangeConsumeURLValues(store, amount, verifyCode)
		sms = conf.GetConsumeSMS()
	case 3:
		// 消费成功通知
		values = exchangeConsumeNotifyURLValues(store, amount, balance)
		sms = conf.GetConsumeNotifySMS()
	case 4:
		// 转卡成功通知
		values = exchangeTransferURLValues(amount)
		sms = conf.GetTransferSMS()
	}
	// return verifyCode, nil
	param := url.Values{"apikey": {sms.ApiKey}, "mobile": {mobile}, "tpl_id": {sms.TemplateID}, "tpl_value": {values.Encode()}}
	resp, err := http.PostForm(sms.YunpianURL, param)
	if err != nil {
		// return "", errkit.Wrap(err, "发送短信失败")
		return "", errkit.New("手机号有误，请重新输入")

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// return "", errkit.Wrap(err, "发送短信失败")
		return "", errkit.New("手机号有误，请重新输入")

	}
	result := &struct {
		Code   int    `json:"code"`
		Msg    string `json:"msg"`
		Result struct {
			Count int     `json:"count"`
			Fee   float32 `json:"fee"`
			Sid   int64   `json:"sid"`
		}
		HTTPStatusCode string `json:"http_status_code"`
		Detail         string `json:"detail"`
	}{}
	err = json.Unmarshal(body, result)
	if err != nil {
		// return "", errkit.Wrap(err, "发送短信失败")
		return "", errkit.New("手机号有误，请重新输入")
	}
	if result.Code != 0 {
		log.Println("++++++++++++result.Code++++++++++++++++++")
		log.Println(result.Code)
		log.Println(result.Msg)
		switch result.Code {
		case 33:
			return "", errkit.Wrap(errkit.New("请求失败"), "短信验证码5分钟内有效，请耐心等待")
		case 53:
			return "", errkit.Wrap(errkit.New("请求失败"), "达到每小时发送验证码数量上限，请一个小时后重试")
		default:
			return "", errkit.Wrap(errkit.New("请求失败"), result.Msg)
		}
	}
	log.Println("++++++++++++result++++++++++++++++++")
	log.Println(result)
	return verifyCode, nil
}
