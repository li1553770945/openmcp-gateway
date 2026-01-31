package auth

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-jwt/jwt/v5"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	user_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/user"
	"github.com/li1553770945/openmcp-gateway/biz/model/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	Repo   user_repo.IUserRepository
	JWTKey string
}

func NewAuthService(userRepo user_repo.IUserRepository, cfg *config.Config) IAuthService {
	return &AuthServiceImpl{
		Repo:   userRepo,
		JWTKey: cfg.AuthConfig.JWTKey, // In real app, load from config
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	hlog.CtxInfof(ctx, "收到用户 %s 的登录请求", req.Username)

	userEntity, err := s.checkUsernameAndPasswd(ctx, req.Username, req.Password)
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

func (s *AuthServiceImpl) checkUsernameAndPasswd(ctx context.Context, username string, password string) (*domain.UserEntity, error) {
	user, err := s.Repo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
