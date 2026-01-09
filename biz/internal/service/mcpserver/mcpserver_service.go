package mcpserver

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"github.com/li1553770945/openmcp-gateway/biz/infra/cache"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	mcpserver_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/mcpserver"
	"github.com/li1553770945/openmcp-gateway/biz/model/mcpserver"
)

type MCPServerServiceImpl struct {
	Repo       mcpserver_repo.IMCPServerRepository
	ProxyCache cache.IProxyCache
}

func NewMCPServerService(repo mcpserver_repo.IMCPServerRepository, proxyCache cache.IProxyCache) IMCPServerService {
	return &MCPServerServiceImpl{
		Repo:       repo,
		ProxyCache: proxyCache,
	}
}

func (s *MCPServerServiceImpl) AddMCPServer(ctx context.Context, req *mcpserver.AddMCPServerReq) (resp *mcpserver.AddMCPServerResp) {
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		return &mcpserver.AddMCPServerResp{
			Code:    constant.InvalidInput,
			Message: "URL 必须以 http:// 或 https:// 开头",
		}
	}

	creatorID, _ := ctx.Value("user_id").(int64)

	entity := &domain.MCPServerEntity{
		Name:        req.Name,
		Description: req.Description,
		Url:         req.URL,
		IsPublic:    req.IsPublic,
		OpenProxy:   req.OpenProxy,
		CreatorID:   creatorID,
	}

	err := s.Repo.SaveMCPServer(entity)
	if err != nil {
		hlog.Errorf("创建 MCPServer 失败: %v", err)
		return &mcpserver.AddMCPServerResp{
			Code:    constant.SystemError,
			Message: "创建 MCPServer 失败",
		}
	}

	return &mcpserver.AddMCPServerResp{
		Code:    constant.Success,
		Message: "成功",
	}
}

func (s *MCPServerServiceImpl) GenerateToken(ctx context.Context, req *mcpserver.GenerateTokenReq) (resp *mcpserver.GenerateTokenResp) {
	// First check if server exists
	server, err := s.Repo.FindMCPServerById(req.ID)
	if err != nil {
		return &mcpserver.GenerateTokenResp{Code: constant.SystemError, Message: "数据库错误"}
	}
	if server == nil {
		return &mcpserver.GenerateTokenResp{Code: constant.NotFound, Message: "McpServer 未找到"}
	}

	// Permission check
	creatorID, _ := ctx.Value("user_id").(int64)
	if server.CreatorID != creatorID {
		return &mcpserver.GenerateTokenResp{Code: constant.Unauthorized, Message: "您无权限生成该 MCPServer 的 Token"}
	}

	tokenStr := generateRandomToken(32)
	retryCount := 0
	for {
		if retryCount >= 10 {
			return &mcpserver.GenerateTokenResp{Code: constant.SystemError, Message: "生成唯一 token 失败，请稍后重试"}
		}
		exist, err := s.Repo.FindTokenByToken(tokenStr)
		if err != nil {
			return &mcpserver.GenerateTokenResp{Code: constant.SystemError, Message: "生成 token 校验失败"}
		}
		if exist == nil {
			break
		}
		tokenStr = generateRandomToken(32)
		retryCount++
	}

	tokenEntity := &domain.MCPServerTokenEntity{
		MCPServerID: req.ID,
		Description: req.Description,
		Token:       tokenStr,
	}

	err = s.Repo.SaveToken(tokenEntity)
	if err != nil {
		return &mcpserver.GenerateTokenResp{Code: constant.SystemError, Message: "保存 token 失败"}
	}

	return &mcpserver.GenerateTokenResp{
		Code:    constant.Success,
		Message: "成功",
		Data:    TokenEntityToRespData(tokenEntity),
	}
}

func (s *MCPServerServiceImpl) GetSelfMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp) {
	creatorID, _ := ctx.Value("user_id").(int64)

	list, err := s.Repo.ListMCPServersByCreatorId(creatorID, req.Start, req.End)
	if err != nil {
		return &mcpserver.GetMCPServerListResp{Code: constant.SystemError, Message: "获取服务器列表失败"}
	}

	data := make([]*mcpserver.GetMCPServerListRespData, 0)
	for _, v := range list {
		data = append(data, EntityToMCPServerListRespData(v))
	}

	return &mcpserver.GetMCPServerListResp{
		Code:    constant.Success,
		Message: "成功",
		Data:    data,
	}
}

func (s *MCPServerServiceImpl) GetPublicMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp) {
	list, err := s.Repo.ListPublicMCPServers(req.Start, req.End)
	if err != nil {
		return &mcpserver.GetMCPServerListResp{Code: constant.SystemError, Message: "获取公开服务器列表失败"}
	}

	data := make([]*mcpserver.GetMCPServerListRespData, 0)
	for _, v := range list {
		data = append(data, EntityToMCPServerListRespData(v))
	}

	return &mcpserver.GetMCPServerListResp{
		Code:    constant.Success,
		Message: "成功",
		Data:    data,
	}
}

func (s *MCPServerServiceImpl) UpdateMCPServer(ctx context.Context, req *mcpserver.UpdateMCPServerReq) (resp *mcpserver.UpdateMCPServerResp) {
	server, err := s.Repo.FindMCPServerById(req.ID)
	if server == nil {
		return &mcpserver.UpdateMCPServerResp{Code: constant.NotFound, Message: "McpServer 未找到"}
	}

	creatorID, _ := ctx.Value("user_id").(int64)
	if server.CreatorID != creatorID {
		return &mcpserver.UpdateMCPServerResp{Code: constant.Unauthorized, Message: "您无权限更新该 MCPServer"}
	}

	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		return &mcpserver.UpdateMCPServerResp{
			Code:    constant.InvalidInput,
			Message: "URL 必须以 http:// 或 https:// 开头",
		}
	}
	oldURL := server.Url
	newURL := req.URL

	server.Name = req.Name
	server.Description = req.Description
	server.Url = req.URL
	server.IsPublic = req.IsPublic
	server.OpenProxy = req.OpenProxy

	err = s.Repo.SaveMCPServer(server)
	if err != nil {
		return &mcpserver.UpdateMCPServerResp{Code: constant.SystemError, Message: "更新服务器失败"}
	}

	if oldURL != newURL {
		tokens, err := s.Repo.FindTokensByMcpServerId(server.ID)
		if err != nil {
			hlog.Errorf("UpdateMCPServer: 通知 ProxyService 清理缓存失败: %v", err)
			return &mcpserver.UpdateMCPServerResp{Code: constant.SystemError, Message: "更新服务器成功但缓存失败"}

		}
		// URL 变更，通知 ProxyService 清理缓存
		for _, t := range tokens {
			err = s.ProxyCache.InvalidateByToken(t.Token)
			if err != nil {
				hlog.Errorf("UpdateMCPServer: 通知 ProxyService 清理缓存失败: %v", err)
				return &mcpserver.UpdateMCPServerResp{Code: constant.SystemError, Message: "更新服务器成功但缓存失败"}
			}
		}

	}
	return &mcpserver.UpdateMCPServerResp{Code: constant.Success, Message: "成功"}
}

func (s *MCPServerServiceImpl) GetMCPServer(ctx context.Context, req *mcpserver.GetMCPServerReq) (resp *mcpserver.GetMCPServerResp) {
	server, err := s.Repo.FindMCPServerById(req.McpServerId)
	if err != nil {
		return &mcpserver.GetMCPServerResp{Code: constant.SystemError, Message: "获取服务器详情失败"}
	}
	if server == nil {
		return &mcpserver.GetMCPServerResp{Code: constant.NotFound, Message: "McpServer 未找到"}
	}

	creatorID, _ := ctx.Value("user_id").(int64)
	if !server.IsPublic && server.CreatorID != creatorID {
		return &mcpserver.GetMCPServerResp{Code: constant.Unauthorized, Message: "您无权限查看该 MCPServer 详情"}
	}

	return &mcpserver.GetMCPServerResp{
		Code:    constant.Success,
		Message: "成功",
		Data:    EntityToMCPServerRespData(server),
	}
}

func (s *MCPServerServiceImpl) DeleteMCPServer(ctx context.Context, req *mcpserver.DeleteMCPServerReq) (resp *mcpserver.DeleteMCPServerResp) {
	server, err := s.Repo.FindMCPServerById(req.ID)
	if err != nil {
		return &mcpserver.DeleteMCPServerResp{Code: constant.SystemError, Message: "查询服务器失败"}
	}
	if server == nil {
		return &mcpserver.DeleteMCPServerResp{Code: constant.NotFound, Message: "McpServer 未找到"}
	}

	creatorID, _ := ctx.Value("user_id").(int64)
	if server.CreatorID != creatorID {
		return &mcpserver.DeleteMCPServerResp{Code: constant.Unauthorized, Message: "您无权限删除该 MCPServer"}
	}
	tokens, err := s.Repo.FindTokensByMcpServerId(server.ID)
	if err != nil {
		hlog.Errorf("DeleteMCPServer: 查询关联 Tokens 失败: %v", err)
		return &mcpserver.DeleteMCPServerResp{Code: constant.SystemError, Message: "删除服务器失败"}
	}
	err = s.Repo.DeleteMCPServer(req.ID)
	if err != nil {
		hlog.Errorf("DeleteMCPServer: 删除 MCPServer 失败: %v", err)
		return &mcpserver.DeleteMCPServerResp{Code: constant.SystemError, Message: "删除服务器失败"}
	}

	// Server 删除后，清理相关缓存
	if s.ProxyCache != nil {
		// MCPServer，通知 ProxyService 清理缓存
		for _, t := range tokens {
			err = s.ProxyCache.InvalidateByToken(t.Token)
			if err != nil {
				hlog.Errorf("DeleteMCPServer: 通知 ProxyService 清理缓存失败: %v", err)
				return &mcpserver.DeleteMCPServerResp{Code: constant.SystemError, Message: "删除服务器成功但缓存清理失败"}
			}
		}
	}

	return &mcpserver.DeleteMCPServerResp{Code: constant.Success, Message: "删除成功"}
}

func (s *MCPServerServiceImpl) DeleteToken(ctx context.Context, req *mcpserver.DeleteTokenReq) (resp *mcpserver.DeleteTokenResp) {
	token, err := s.Repo.FindTokenById(req.ID)
	if err != nil {
		return &mcpserver.DeleteTokenResp{Code: constant.SystemError, Message: "查询Token失败"}
	}
	if token == nil {
		return &mcpserver.DeleteTokenResp{Code: constant.NotFound, Message: "Token 未找到"}
	}

	// 校验权限
	server, err := s.Repo.FindMCPServerById(token.MCPServerID)
	if err != nil {
		return &mcpserver.DeleteTokenResp{Code: constant.SystemError, Message: "查询服务器失败"}
	}
	if server == nil {
		return &mcpserver.DeleteTokenResp{Code: constant.NotFound, Message: "关联的服务器未找到"}
	}

	creatorID, _ := ctx.Value("user_id").(int64)
	if server.CreatorID != creatorID {
		return &mcpserver.DeleteTokenResp{Code: constant.Unauthorized, Message: "您无权限删除该 Token"}
	}

	err = s.Repo.DeleteToken(req.ID)
	if err != nil {
		return &mcpserver.DeleteTokenResp{Code: constant.SystemError, Message: "删除Token失败"}
	}

	// 清除 Token 缓存
	if s.ProxyCache != nil {
		s.ProxyCache.InvalidateByToken(token.Token)
	}

	return &mcpserver.DeleteTokenResp{Code: constant.Success, Message: "删除成功"}
}

func generateRandomToken(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
