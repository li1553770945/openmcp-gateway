package mcpserver

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/model/mcpserver"
)

type IMCPServerService interface {
	AddMCPServer(ctx context.Context, req *mcpserver.AddMCPServerReq) (resp *mcpserver.AddMCPServerResp)
	GenerateToken(ctx context.Context, req *mcpserver.GenerateTokenReq) (resp *mcpserver.GenerateTokenResp)
	GetSelfMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp)
	GetPublicMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp)
	UpdateMCPServer(ctx context.Context, req *mcpserver.UpdateMCPServerReq) (resp *mcpserver.UpdateMCPServerResp)
	GetMCPServer(ctx context.Context, req *mcpserver.GetMCPServerReq) (resp *mcpserver.GetMCPServerResp)
}
