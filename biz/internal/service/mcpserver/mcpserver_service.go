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
		hlog.Errorf("Create MCPServer error: %v", err)
		return &mcpserver.AddMCPServerResp{
			Code:    constant.SystemError,
			Message: "Failed to create MCPServer",
		}
	}

	return &mcpserver.AddMCPServerResp{
		Code:    constant.Success,
		Message: "Success",
	}
}

func (s *MCPServerServiceImpl) GenerateToken(ctx context.Context, req *mcpserver.GenerateTokenReq) (resp *mcpserver.GenerateTokenResp) {
	// First check if server exists
	server, err := s.Repo.FindMCPServerById(req.ID)
	if err != nil {
		return &mcpserver.GenerateTokenResp{Code: constant.SystemError, Message: "DB Error"}
	}
	if server == nil {
		return &mcpserver.GenerateTokenResp{Code: constant.NotFound, Message: "Server not found"}
	}

	// Permission check
	creatorID, _ := ctx.Value("user_id").(int64)
	if server.CreatorID != creatorID {
		return &mcpserver.GenerateTokenResp{Code: constant.Unauthorized, Message: "Unauthorized"}
	}

	tokenStr := generateRandomToken(32)
	tokenEntity := &domain.MCPServerTokenEntity{
		MCPServerID: req.ID,
		Description: req.Description,
		Token:       tokenStr,
	}

	err = s.Repo.SaveToken(tokenEntity)
	if err != nil {
		return &mcpserver.GenerateTokenResp{Code: constant.SystemError, Message: "Failed to save token"}
	}

	return &mcpserver.GenerateTokenResp{
		Code:    constant.Success,
		Message: "Success",
		Data:    TokenEntityToRespData(tokenEntity),
	}
}

func (s *MCPServerServiceImpl) GetSelfMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp) {
	creatorID, _ := ctx.Value("user_id").(int64)

	list, err := s.Repo.ListMCPServersByCreatorId(creatorID, req.Start, req.End)
	if err != nil {
		return &mcpserver.GetMCPServerListResp{Code: constant.SystemError, Message: err.Error()}
	}

	data := make([]*mcpserver.GetMCPServerListRespData, len(list))
	for i, v := range list {
		data[i] = EntityToMCPServerListRespData(v)
	}

	return &mcpserver.GetMCPServerListResp{
		Code:    constant.Success,
		Message: "Success",
		Data:    data,
	}
}

func (s *MCPServerServiceImpl) GetPublicMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp) {
	list, err := s.Repo.ListPublicMCPServers(req.Start, req.End)
	if err != nil {
		return &mcpserver.GetMCPServerListResp{Code: constant.SystemError, Message: err.Error()}
	}

	data := make([]*mcpserver.GetMCPServerListRespData, len(list))
	for i, v := range list {
		data[i] = EntityToMCPServerListRespData(v)
	}

	return &mcpserver.GetMCPServerListResp{
		Code:    constant.Success,
		Message: "Success",
		Data:    data,
	}
}

func (s *MCPServerServiceImpl) UpdateMCPServer(ctx context.Context, req *mcpserver.UpdateMCPServerReq) (resp *mcpserver.UpdateMCPServerResp) {
	server, err := s.Repo.FindMCPServerById(req.ID)
	if server == nil {
		return &mcpserver.UpdateMCPServerResp{Code: constant.NotFound, Message: "Not Found"}
	}

	creatorID, _ := ctx.Value("user_id").(int64)
	if server.CreatorID != creatorID {
		return &mcpserver.UpdateMCPServerResp{Code: constant.Unauthorized, Message: "Unauthorized"}
	}

	server.Name = req.Name
	server.Description = req.Description
	server.Url = req.URL
	server.IsPublic = req.IsPublic
	server.OpenProxy = req.OpenProxy

	err = s.Repo.SaveMCPServer(server)
	if err != nil {
		return &mcpserver.UpdateMCPServerResp{Code: constant.SystemError, Message: err.Error()}
	}

	return &mcpserver.UpdateMCPServerResp{Code: constant.Success, Message: "Success"}
}

func (s *MCPServerServiceImpl) GetMCPServer(ctx context.Context, req *mcpserver.GetMCPServerReq) (resp *mcpserver.GetMCPServerResp) {
	server, err := s.Repo.FindMCPServerById(req.McpServerId)
	if err != nil {
		return &mcpserver.GetMCPServerResp{Code: constant.SystemError, Message: err.Error()}
	}
	if server == nil {
		return &mcpserver.GetMCPServerResp{Code: constant.NotFound, Message: "Not Found"}
	}

	creatorID, _ := ctx.Value("user_id").(int64)
	if !server.IsPublic && server.CreatorID != creatorID {
		return &mcpserver.GetMCPServerResp{Code: constant.Unauthorized, Message: "Access Denied"}
	}

	return &mcpserver.GetMCPServerResp{
		Code:    constant.Success,
		Message: "Success",
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
