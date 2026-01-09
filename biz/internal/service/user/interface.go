package user

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/model/user"
)

type IUserService interface {
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp)
	GetSelfInfo(ctx context.Context) (resp *user.GetUserInfoResp)
	UpdateSelfInfo(ctx context.Context, req *user.UpdateSelfInfoReq) (resp *user.UpdateSelfInfoResp)
	Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error)
}
