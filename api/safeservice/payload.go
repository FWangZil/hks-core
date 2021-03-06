package safeservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FWangZil/errkit"
	"github.com/gin-gonic/gin"
)

const (
	respOk      = 0 //"OK"
	respNotUser = 1 //"NotUser"
	respFail    = 2 //"FAIL"
	respUnLogin = 3 //"UnLogin"
	respNoAuth  = 4 //"NoAuth"
)

// resp Payload返回
type resp map[string]interface{}

func ok(c *gin.Context, resp resp) {
	result := make(map[string]interface{})
	if resp != nil {
		for key, value := range resp {
			if fmt.Sprint(value) != "<nil>" {
				result[key] = value
			}
		}
	}
	result["resultCode"] = respOk
	result["resultMsg"] = "成功"

	c.JSON(http.StatusOK, result)
}

func okArr(c *gin.Context, resp resp) {
	result := make(map[string]interface{})
	if resp != nil {
		for key, value := range resp {
			if fmt.Sprint(value) != "<nil>" {
				data := make(map[string]interface{})
				data[key] = value
				result[key] = data
			}
		}
	}
	result["resultCode"] = respOk
	result["resultMsg"] = "成功"

	c.JSON(http.StatusOK, result)
}

// unLogin 未登录情况返回
func unLogin(c *gin.Context) {
	c.JSON(http.StatusOK, resp{
		"resultCode": respUnLogin,
		"resultMsg":  "未登录或登录失效",
	})
}

// noAuth 无权限
func noAuth(c *gin.Context) {
	c.JSON(http.StatusOK, resp{
		"resultCode": respNoAuth,
		"resultMsg":  "无权限",
	})
}

// notVIP 返回非会员错误
func notVIP(c *gin.Context, e error) {
	logError(e)
	message := errkit.Msg(e)
	c.JSON(http.StatusOK, resp{
		"resultCode": respNotUser,
		"resultMsg":  message,
	})
}

// fail 错误情况返回
func fail(c *gin.Context, e error) {
	logError(e)
	message := errkit.Msg(e)
	c.JSON(http.StatusOK, resp{
		"resultCode": respFail,
		"resultMsg":  message,
	})
}

func pageOK(c *gin.Context, html string, param gin.H) {
	c.HTML(http.StatusOK, html, param)
}

func pageFail(c *gin.Context, e error) {
	logError(e)
	message := errkit.Msg(e)
	param := gin.H{
		"title":   "出错啦",
		"message": message,
	}
	c.HTML(http.StatusOK, "w_error.html", param)
}

func logError(e error) {
	log.Printf("error: 【full】 %+#v ", e)
}
