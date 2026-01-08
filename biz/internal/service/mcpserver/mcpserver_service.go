package mcpserver

import (
	"context"
	"math/rand"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	mcpserver_repo "github.com/li1553770945/openmcp-gateway/biz/internal/repo/mcpserver"
	"github.com/li1553770945/openmcp-gateway/biz/model/mcpserver"
)

type MCPServerServiceImpl struct {
	Repo mcpserver_repo.IMCPServerRepository
}

func NewMCPServerService(repo mcpserver_repo.IMCPServerRepository) IMCPServerService {
	return &MCPServerServiceImpl{
		Repo: repo,
	}
}

func (s *MCPServerServiceImpl) AddMCPServer(ctx context.Context, req *mcpserver.AddMCPServerReq) (resp *mcpserver.AddMCPServerResp) {
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
		return &mcpserver.GenerateTokenResp{Code: constant.Unauthorized, Message: "无权限"}
	}

	tokenStr := generateRandomToken(32)
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

	data := make([]*mcpserver.GetMCPServerListRespData, len(list))
	for i, v := range list {
		data[i] = EntityToMCPServerListRespData(v)
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

	data := make([]*mcpserver.GetMCPServerListRespData, len(list))
	for i, v := range list {
		data[i] = EntityToMCPServerListRespData(v)
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
		return &mcpserver.UpdateMCPServerResp{Code: constant.Unauthorized, Message: "无权限"}
	}

	server.Name = req.Name
	server.Description = req.Description
	server.Url = req.URL
	server.IsPublic = req.IsPublic
	server.OpenProxy = req.OpenProxy

	err = s.Repo.SaveMCPServer(server)
	if err != nil {
		return &mcpserver.UpdateMCPServerResp{Code: constant.SystemError, Message: "更新服务器失败"}
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
		return &mcpserver.GetMCPServerResp{Code: constant.Unauthorized, Message: "无权限"}
	}

	return &mcpserver.GetMCPServerResp{
		Code:    constant.Success,
		Message: "成功",
		Data:    EntityToMCPServerRespData(server),
	}
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
