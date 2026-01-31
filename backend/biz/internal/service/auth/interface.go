package auth

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"github.com/li1553770945/openmcp-gateway/biz/model/auth"
)

type IAuthService interface {
	Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error)
	checkUsernameAndPasswd(ctx context.Context, username string, password string) (*domain.UserEntity, error)
}
