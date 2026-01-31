package container

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/openmcp-gateway/biz/infra/cache"
	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"github.com/li1553770945/openmcp-gateway/biz/infra/database"
	"github.com/li1553770945/openmcp-gateway/biz/internal/middleware"
	mcpserver_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/mcpserver"
	user_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/user"
	authService "github.com/li1553770945/openmcp-gateway/biz/internal/service/auth"
	mcpserver_service "github.com/li1553770945/openmcp-gateway/biz/internal/service/mcpserver"
	proxy_service "github.com/li1553770945/openmcp-gateway/biz/internal/service/proxy"
	user_service "github.com/li1553770945/openmcp-gateway/biz/internal/service/user"
)

type Container struct {
	UserService      user_service.IUserService
	AuthService      authService.IAuthService
	MCPServerService mcpserver_service.IMCPServerService
	ProxyService     proxy_service.IProxyService

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
	mcpserverRepo := mcpserver_repo.NewMCPServerRepository(db)

	proxyCache := cache.NewProxyCache(&cfg.ProxyConfig)

	userSvc := user_service.NewUserService(userRepo)
	mcpserverSvc := mcpserver_service.NewMCPServerService(mcpserverRepo, proxyCache)
	proxySvc := proxy_service.NewProxyService(mcpserverRepo, proxyCache)

	globalApp = &Container{
		UserService:               userSvc,
		AuthService:               authService.NewAuthService(userRepo, cfg),
		MCPServerService:          mcpserverSvc,
		ProxyService:              proxySvc,
		Config:                    cfg,
		AuthAndUserInfoMiddleware: middleware.NewAuthAndUserInfoMiddleware(cfg),
	}
}
