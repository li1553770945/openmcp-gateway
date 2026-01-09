package assembler

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/do"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
)

func UserDoToEntity(do *do.UserDO) *domain.UserEntity {
	return &domain.UserEntity{
		ID:       do.ID,
		Username: do.Username,
		Nickname: do.Nickname,
		Role:     do.Role,
		CanUse:   do.CanUse,
		Email:    do.Email,
		Password: do.Password,
	}
}

func UserEntityToDo(entity *domain.UserEntity) *do.UserDO {
	userDO := &do.UserDO{
		Username: entity.Username,
		Nickname: entity.Nickname,
		Role:     entity.Role,
		CanUse:   entity.CanUse,
		Email:    entity.Email,
	}
	if entity.Password != "" {
		userDO.Password = entity.Password
	}
	return userDO
}
