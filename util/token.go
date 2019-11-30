package util

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/meikeland/errkit"
	"github.com/dgrijalva/jwt-go"
)

var (
	expires = 30 * 365 * 24 * time.Hour
	// expires     = 1 * time.Minute // 测试
	tokenKey    = "user-token"
	tokenSecret = "3a7f3df7-50b3-4699-abaa-c0d191e18b2d"
)

// Token 登录
type Token struct {
	UserID   uint      `json:"userID"`
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireAt"`
}

// Sign 加密Token
func Sign(userID uint) (*Token, error) {
	now := time.Now()
	expiresAt := now.Add(expires)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", userID),
		ExpiresAt: expiresAt.Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    tokenKey,
	})
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return nil, errkit.Wrap(err, "加密登录失败")
	}
	return &Token{
		UserID:   userID,
		Token:    tokenString,
		ExpireAt: expiresAt,
	}, nil
}

// Unsign 解密Token
func Unsign(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		})
	if err != nil {
		log.Printf("%+v", token)
		log.Printf("该用户的token已失效: %v", err)
	}
	userID, err := strconv.Atoi(token.Claims.(*jwt.StandardClaims).Id)
	if err != nil {
		return 0, errkit.Wrap(err, "解密登录失败")
	}
	return uint(userID), nil
}
