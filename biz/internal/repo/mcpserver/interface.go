package mcpserver

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
)

type IMCPServerRepository interface {
	SaveMCPServer(server *domain.MCPServerEntity) error
	FindMCPServerById(id int64) (*domain.MCPServerEntity, error)
	ListMCPServersByCreatorId(creatorId int64, start, end int64) ([]*domain.MCPServerEntity, error)
	ListPublicMCPServers(start, end int64) ([]*domain.MCPServerEntity, error)

	SaveToken(token *domain.MCPServerTokenEntity) error
	FindTokenByToken(token string) (*domain.MCPServerTokenEntity, error)
	FindTokenById(id int64) (*domain.MCPServerTokenEntity, error)
	FindTokensByMcpServerId(mcpServerId int64) ([]*domain.MCPServerTokenEntity, error)

	DeleteMCPServer(id int64) error
	DeleteToken(id int64) error
}
