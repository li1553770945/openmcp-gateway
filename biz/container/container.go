package container

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"github.com/li1553770945/openmcp-gateway/biz/infra/database"
	"github.com/li1553770945/openmcp-gateway/biz/internal/middleware"
	user_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/user"
	"github.com/li1553770945/openmcp-gateway/biz/internal/service/user"
	userService "github.com/li1553770945/openmcp-gateway/biz/internal/service/user"
)

type Container struct {
	UserService user.IUserService

	Config                    *config.Config
	AuthAndUserInfoMiddleware app.HandlerFunc
}

var globalApp *Container

func GetGlobalContainer() *Container {
	if globalApp == nil {
		panic("全局app未初始化")
	}
	return globalApp
}

func InitContainer(env string) {
	cfg := config.GetConfig(env)
	db := database.NewDatabase(cfg)
	userRepo := user_repo.NewUserRepository(db)

	globalApp = &Container{
		UserService:               userService.NewUserService(userRepo),
		Config:                    cfg,
		AuthAndUserInfoMiddleware: middleware.NewAuthAndUserInfoMiddleware(),
	}
}
