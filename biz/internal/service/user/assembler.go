package user

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"github.com/li1553770945/openmcp-gateway/biz/model/user"
)

func EntityToUserInfoData(entity *domain.UserEntity) *user.GetUserInfoRespData {
	return &user.GetUserInfoRespData{
		Username: entity.Username,
		Nickname: entity.Nickname,
		Role:     entity.Role,
	}
}
