package user

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
)

type IUserRepository interface {
	FindUserByUsername(username string) (*domain.UserEntity, error)
	FindUserById(userId int64) (*domain.UserEntity, error)
	SaveUser(user *domain.UserEntity) error
}
