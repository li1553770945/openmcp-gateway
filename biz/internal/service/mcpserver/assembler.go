package mcpserver

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"github.com/li1553770945/openmcp-gateway/biz/model/mcpserver"
)

func EntityToMCPServerListRespData(entity *domain.MCPServerEntity) *mcpserver.GetMCPServerListRespData {
	if entity == nil {
		return nil
	}
	return &mcpserver.GetMCPServerListRespData{
		Name:        entity.Name,
		Description: entity.Description,
		URL:         entity.Url,
		IsPublic:    entity.IsPublic,
		OpenProxy:   entity.OpenProxy,
	}
}

func EntityToMCPServerRespData(entity *domain.MCPServerEntity) *mcpserver.GetMCPServerRespData {
	if entity == nil {
		return nil
	}
	tokens := make([]*mcpserver.TokenData, 0, len(entity.Tokens))
	for _, t := range entity.Tokens {
		tokens = append(tokens, &mcpserver.TokenData{
			Token:       t.Token,
			Description: t.Description,
		})
	}

	return &mcpserver.GetMCPServerRespData{
		Name:        entity.Name,
		Description: entity.Description,
		URL:         entity.Url,
		IsPublic:    entity.IsPublic,
		OpenProxy:   entity.OpenProxy,
		Token:       tokens,
	}
}

func TokenEntityToRespData(entity *domain.MCPServerTokenEntity) *mcpserver.GenerateTokenRespData {
	if entity == nil {
		return nil
	}
	return &mcpserver.GenerateTokenRespData{
		Token: entity.Token,
	}
}
