package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
)

func NewAuthAndUserInfoMiddleware() app.HandlerFunc {
	return AuthMiddleware()
}

// AuthMiddleware 认证中间件
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取 Authorization 头部的 token
		token := string(c.GetHeader("Authorization"))
		if token == "" {
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "您还未登陆，请先登录"})
			c.Abort()
			return
		}
		// TODO: 为真正实现逻辑
		c.Next(ctx)
	}
}
