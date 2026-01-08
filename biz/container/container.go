package container

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"github.com/li1553770945/openmcp-gateway/biz/infra/database"
	"github.com/li1553770945/openmcp-gateway/biz/internal/middleware"
	user_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/user"
	authService "github.com/li1553770945/openmcp-gateway/biz/internal/service/auth"
	userService "github.com/li1553770945/openmcp-gateway/biz/internal/service/user"
)

type Container struct {
	UserService userService.IUserService
	AuthService authService.IAuthService

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
	userSvc := userService.NewUserService(userRepo)

	globalApp = &Container{
		UserService:               userSvc,
		AuthService:               authService.NewAuthService(userSvc, cfg),
		Config:                    cfg,
		AuthAndUserInfoMiddleware: middleware.NewAuthAndUserInfoMiddleware(),
	}
}
