package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	userRepo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/user"
	"github.com/li1553770945/openmcp-gateway/biz/model/user"
)

type UserServiceImpl struct {
	Repo userRepo.IUserRepository
}

func NewUserService(repo userRepo.IUserRepository) IUserService {
	return &UserServiceImpl{
		Repo: repo,
	}
}

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp) {
	findUser, err := s.Repo.FindUserById(req.UserId)
	if err != nil {
		hlog.Errorf("查询用户信息错误:%s", err.Error())
		resp = &user.GetUserInfoResp{
			Code:    constant.SystemError,
			Message: "系统错误，查询用户信息失败",
		}
		return
	}
	resp = &user.GetUserInfoResp{
		Code: constant.Success,
		Data: EntityToUserInfoData(findUser),
	}
	return
}
