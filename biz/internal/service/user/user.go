package user

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	userRepo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/user"
	"github.com/li1553770945/openmcp-gateway/biz/model/user"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserServiceImpl) GetSelfInfo(ctx context.Context) (resp *user.GetUserInfoResp) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		resp = &user.GetUserInfoResp{
			Code:    constant.Unauthorized,
			Message: "未登录或登录状态已过期",
		}
		return
	}
	findUser, err := s.Repo.FindUserById(userID)
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

func (s *UserServiceImpl) UpdateSelfInfo(ctx context.Context, req *user.UpdateSelfInfoReq) (resp *user.UpdateSelfInfoResp) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		resp = &user.UpdateSelfInfoResp{
			Code:    constant.Unauthorized,
			Message: "未登录或登录状态已过期",
		}
		return
	}

	findUser, err := s.Repo.FindUserById(userID)
	if err != nil {
		hlog.Errorf("查询用户信息错误:%s", err.Error())
		resp = &user.UpdateSelfInfoResp{
			Code:    constant.SystemError,
			Message: "系统错误，查询用户信息失败",
		}
		return
	}

	if findUser == nil {
		resp = &user.UpdateSelfInfoResp{
			Code:    constant.NotFound,
			Message: "用户不存在",
		}
		return
	}

	// update info
	if req.Nickname != nil {
		findUser.Nickname = *req.Nickname
	}
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			hlog.Errorf("密码加密错误:%s", err.Error())
			resp = &user.UpdateSelfInfoResp{
				Code:    constant.SystemError,
				Message: "系统错误，密码处理失败",
			}
			return
		}
		findUser.Password = string(hashedPassword)
	}

	err = s.Repo.SaveUser(findUser)
	if err != nil {
		hlog.Errorf("更新用户信息错误:%s", err.Error())
		resp = &user.UpdateSelfInfoResp{
			Code:    constant.SystemError,
			Message: "系统错误，更新用户信息失败",
		}
		return
	}

	resp = &user.UpdateSelfInfoResp{
		Code:    constant.Success,
		Message: "更新成功",
	}
	return
}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	hlog.CtxInfof(ctx, "收到用户 %s 的注册请求", req.Username)

	existingUser, _ := s.Repo.FindUserByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &domain.UserEntity{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     "user",
		CanUse:   true,
		Email:    req.Email,
		Nickname: req.Nickname,
	}

	err = s.Repo.SaveUser(newUser)
	if err != nil {
		hlog.CtxErrorf(ctx, "注册失败: %v", err)
		return &user.RegisterResp{
			Code:    constant.InvalidInput, // Or duplicate user code
			Message: err.Error(),
		}, nil
	}

	hlog.CtxInfof(ctx, "注册成功: %v", newUser.ID)
	return &user.RegisterResp{
		Code:    constant.Success,
		Message: "注册成功",
	}, nil
}
