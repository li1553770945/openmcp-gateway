package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
)

func NewAuthAndUserInfoMiddleware(cfg *config.Config) app.HandlerFunc {
	return AuthMiddleware(cfg)
}

// AuthMiddleware 认证中间件
func AuthMiddleware(cfg *config.Config) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取 Authorization 头部的 token
		tokenString := string(c.GetHeader("Authorization"))
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[len("Bearer "):]
		}
		if tokenString == "" {
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "您还未登陆，请先登录"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.AuthConfig.JWTKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "Token无效或已过期"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userIdFloat, ok := claims["userId"].(float64); ok {
				userId := int64(userIdFloat)
				ctx = context.WithValue(ctx, "user_id", userId)
			}
		}

		c.Next(ctx)
	}
}
