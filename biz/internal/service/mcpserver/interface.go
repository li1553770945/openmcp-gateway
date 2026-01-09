package mcpserver

import (
	"context"

	"github.com/li1553770945/openmcp-gateway/biz/model/mcpserver"
)

type IMCPServerService interface {
	AddMCPServer(ctx context.Context, req *mcpserver.AddMCPServerReq) (resp *mcpserver.AddMCPServerResp)
	GenerateToken(ctx context.Context, req *mcpserver.GenerateTokenReq) (resp *mcpserver.GenerateTokenResp)
	GetMCPServerList(ctx context.Context, req *mcpserver.GetMCPServerListReq) (resp *mcpserver.GetMCPServerListResp)
	GetMCPServerCount(ctx context.Context, req *mcpserver.GetMCPServerCountReq) (resp *mcpserver.GetMCPServerCountResp)
	UpdateMCPServer(ctx context.Context, req *mcpserver.UpdateMCPServerReq) (resp *mcpserver.UpdateMCPServerResp)
	GetMCPServer(ctx context.Context, req *mcpserver.GetMCPServerReq) (resp *mcpserver.GetMCPServerResp)
	DeleteMCPServer(ctx context.Context, req *mcpserver.DeleteMCPServerReq) (resp *mcpserver.DeleteMCPServerResp)
	DeleteToken(ctx context.Context, req *mcpserver.DeleteTokenReq) (resp *mcpserver.DeleteTokenResp)
}
