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

func (s *UserServiceImpl) CheckUsernameAndPasswd(ctx context.Context, username string, password string) (*domain.UserEntity, error) {
	user, err := s.Repo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password mismatch")
	}
	return user, nil
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, username string, password string, email string) (*domain.UserEntity, error) {
	// Check if user exists
	existingUser, _ := s.Repo.FindUserByUsername(username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &domain.UserEntity{
		Username: username,
		Password: string(hashedPassword),
		Role:     "user",
		CanUse:   true,
		Email:    email,
	}

	err = s.Repo.SaveUser(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
