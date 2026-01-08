package user

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/model/user"
)

type IUserService interface {
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp)
}
