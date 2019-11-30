package safeservice

import (
	"fmt"
	"hks/hks-core/internal/user"
	"hks/hks-core/pkg"
	"hks/hks-core/util"
	"log"

	"github.com/FWangZil/errkit"
	"github.com/gin-gonic/gin"
)

// getUserByQuery 通过ID获取用户信息
func getUserByQuery(c *gin.Context) {
	userID, err := util.ParseUint(c.Query("userID"))
	mobile := c.Query("mobile")
	if err != nil {
		fail(c, err)
		return
	}
	query := user.Query{
		UserID: userID,
		Mobile: mobile,
	}
	userInfo, err := user.Repo.GetUserByQuery(query)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": userInfo,
	})
}

// listUser 获取用户列表信息
func listUser(c *gin.Context) {
	param := struct {
		UserID uint `json:"userID" form:"userID"` // 用户自增ID
		Gender uint `json:"gender" form:"gender"` // 性别
		Type   uint `json:"type" form:"type"`     // 用户类型
	}{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	pagination := paginationByGet(c)
	query := user.Query{
		UserID: param.UserID,
		Limit:  pagination.PageSize,
		Offset: pagination.Offset,
		Gender: param.Gender,
		Type:   param.Type,
	}
	users, count, err := user.Repo.ListUser(query)
	if err != nil {
		fail(c, err)
		return
	}
	pagination.Total = count
	ok(c, resp{
		"data": users,
	})
}

// registerUser 用户注册接口
func registerUser(c *gin.Context) {
	param := pkg.User{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	if len(param.Password) < 6 {
		fail(c, fmt.Errorf("密码不能小于六位"))
		return
	}
	if len(param.Mobile) > 0 {
		if err := util.ValiateMobileNumber(param.Mobile); err != nil {
			fail(c, fmt.Errorf("手机号码格式错误"))
			return
		}
	}
	if len(param.IDCard) > 0 {
		if err := util.CheckIDCard(param.IDCard); err != nil {
			fail(c, fmt.Errorf("身份号码格式错误"))
			return
		}
	}
	userInfo, err := user.Repo.UserRegister(&param)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": userInfo,
	})
}

// addRelationship 用户增加亲友关系接口
func addRelationship(c *gin.Context) {
	param := pkg.Relative{}
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Println(fmt.Errorf("参数错误:%w", err))
		fail(c, errkit.New("参数错误"))
		return
	}
	if len(param.Mobile) < 0 {
		if err := util.ValiateMobileNumber(param.Mobile); err != nil {
			fail(c, fmt.Errorf("手机号码格式错误"))
			return
		}
	}
	userInfo, err := user.Repo.AddRelationship(&param)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": userInfo,
	})
}

// getRelationship 获取用户亲友关系接口
func getRelationship(c *gin.Context) {
	userID, err := util.ParseUint(c.Query("userID"))
	if userID == 0 {
		fail(c, fmt.Errorf("请传入用户ID"))
		return
	}
	relations, err := user.Repo.GetUserRelationship(userID)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"data": relations,
	})
}
