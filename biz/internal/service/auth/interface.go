package auth

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/model/auth"
)

type IAuthService interface {
	Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error)
	Register(ctx context.Context, req *auth.RegisterReq) (*auth.RegisterResp, error)
}
