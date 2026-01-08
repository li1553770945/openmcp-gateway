package auth

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-jwt/jwt/v5"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"github.com/li1553770945/openmcp-gateway/biz/internal/service/user"
	"github.com/li1553770945/openmcp-gateway/biz/model/auth"
)

type AuthServiceImpl struct {
	UserService user.IUserService
	JWTKey      string
}

func NewAuthService(userService user.IUserService, cfg *config.Config) IAuthService {
	return &AuthServiceImpl{
		UserService: userService,
		JWTKey:      cfg.AuthConfig.JWTKey, // In real app, load from config
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	hlog.CtxInfof(ctx, "收到用户 %s 的登录请求", req.Username)

	userEntity, err := s.UserService.CheckUsernameAndPasswd(ctx, req.Username, req.Password)
	if err != nil {
		hlog.CtxErrorf(ctx, " %s 用户名密码验证失败： 错误: %v", req.Username, err)
		return &auth.LoginResp{
			Code:    constant.Unauthorized,
			Message: "用户名密码验证失败",
		}, nil
	}

	hlog.CtxInfof(ctx, "用户 %s 登录验证成功", req.Username)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userEntity.ID,
			"exp":    time.Now().Add(14 * 24 * time.Hour).Unix(),
		})
	token, err := t.SignedString([]byte(s.JWTKey))
	if err != nil {
		hlog.CtxErrorf(ctx, "jwt 加密失败："+err.Error())
		return &auth.LoginResp{
			Code:    constant.SystemError,
			Message: "系统错误",
		}, err
	}

	return &auth.LoginResp{
		Code:    constant.Success,
		Message: "登录成功",
		Data: &auth.LoginRespData{
			Token: token,
		},
	}, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterReq) (*auth.RegisterResp, error) {
	hlog.CtxInfof(ctx, "收到用户 %s 的注册请求", req.Username)

	userEntity, err := s.UserService.RegisterUser(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		hlog.CtxErrorf(ctx, "注册失败: %v", err)
		return &auth.RegisterResp{
			Code:    constant.InvalidInput, // Or duplicate user code
			Message: err.Error(),
		}, nil
	}

	hlog.CtxInfof(ctx, "注册成功: %v", userEntity.ID)
	return &auth.RegisterResp{
		Code:    constant.Success,
		Message: "注册成功",
	}, nil
}
