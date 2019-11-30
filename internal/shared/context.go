package shared

import (
	"context"
	"hks/hks-core/pkg"
)

type key string

const userKey key = "user"
const adminKey key = "admin"

// WithUser 往已有的context中注入用户信息
func WithUser(ctx context.Context, user *pkg.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser 从context中提取用户信息
func GetUser(ctx context.Context) (*pkg.User, bool) {
	user, ok := ctx.Value(userKey).(*pkg.User)
	return user, ok
}

// WithAdmin 往已有的context中注入管理员信息
func WithAdmin(ctx context.Context, user *pkg.Admin) context.Context {
	return context.WithValue(ctx, adminKey, user)
}

// GetAdmin 从context中提取管理员信息
func GetAdmin(ctx context.Context) (*pkg.Admin, bool) {
	user, ok := ctx.Value(adminKey).(*pkg.Admin)
	return user, ok
}
