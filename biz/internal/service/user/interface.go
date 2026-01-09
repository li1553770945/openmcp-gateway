package user

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"github.com/li1553770945/openmcp-gateway/biz/model/user"
)

type IUserService interface {
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp)
	GetSelfInfo(ctx context.Context) (resp *user.GetUserInfoResp)
	CheckUsernameAndPasswd(ctx context.Context, username string, password string) (*domain.UserEntity, error)
	RegisterUser(ctx context.Context, username string, password string, email string) (*domain.UserEntity, error)
}
