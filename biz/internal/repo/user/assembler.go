package user

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/do"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
)

func DoToEntity(do *do.UserDO) *domain.UserEntity {
	return &domain.UserEntity{
		ID:       do.ID,
		Username: do.Username,
		Nickname: do.Nickname,
		Role:     do.Role,
		CanUse:   do.CanUse,
	}
}

func EntityToDo(entity *domain.UserEntity) *do.UserDO {
	userDO := &do.UserDO{
		Username: entity.Username,
		Nickname: entity.Nickname,
		Role:     entity.Role,
		CanUse:   entity.CanUse,
	}
	if entity.Password != "" {
		userDO.Password = entity.Password
	}
	return userDO
}
