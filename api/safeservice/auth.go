package safeservice

import (
	"hks/hks-core/internal/opauth"
	"hks/hks-core/internal/shared"
	"hks/hks-core/internal/user"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/FWangZil/errkit"
)

// 获取当前登录用户的ID
func getUserIDFromSession(c *gin.Context) uint {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		unLogin(c)
		c.Abort()
	}
	return t.(uint)
}

// auth 拦截浏览器session，判断是否为合法的运营人员，将信息注入到context
func auth(c *gin.Context) {
	userID := getUserIDFromSession(c)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if err != nil {
		unLogin(c)
		c.Abort()
	}
	ctx := shared.WithUser(c.Request.Context(), userInfo)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

// currentUser 获取当前登录用户
func currentUser(c *gin.Context) {
	userInfo, exist := shared.GetUser(c.Request.Context())
	if !exist {
		ok(c, nil)
		return
	}
	session := sessions.Default(c)
	err := session.Save()
	if err != nil {
		log.Println("登录时 session 保存过程出错")
		log.Println(err)
	}
	ok(c, resp{
		"user": userInfo,
	})
}

// currentLoginUser 获取当前登录用户,未登录时不调用
func currentLoginUser(c *gin.Context) {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		ok(c, resp{
			"login": false,
		})
		return
	}
	userID := t.(uint)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if userInfo == nil || err != nil {
		ok(c, nil)
		return
	}
	err = session.Save()
	if err != nil {
		log.Println("登录时 session 保存过程出错")
		log.Println(err)
	}
	ok(c, resp{
		"user":  userInfo,
		"login": true,
	})
}

// login 使用账号和密码进行登录
func login(c *gin.Context) {
	param := &struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBind(param); err != nil {
		fail(c, errkit.Wrap(err, "登录失败，请输入账号和密码"))
		return
	}
	userInfo, err := opauth.ACL.Login(param.Account, param.Password)
	if err != nil {
		fail(c, errkit.New("登录失败，请检查账号和密码是否正确"))
		return
	}
	session := sessions.Default(c)
	session.Set("userID", userInfo.ID)
	err = session.Save()
	if err != nil {
		log.Println(err)
	}
	ok(c, resp{
		"user": userInfo,
	})
}

// logout 退出登录
func logoutUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userID")
	err := session.Save()
	if err != nil {
		log.Println("退出登录时 session 保存过程出错")
		log.Println(err)
	}
	ok(c, nil)
}

// logout 退出登录
func logoutAdmin(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("adminID")
	err := session.Save()
	if err != nil {
		log.Println("退出登录时 session 保存过程出错")
		log.Println(err)
	}
	ok(c, nil)
}
