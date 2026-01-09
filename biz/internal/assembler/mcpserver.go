package assembler

import (
	"time"

	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"github.com/li1553770945/openmcp-gateway/biz/model/mcpserver"
)

func MCPServerEntityToListRespData(entity *domain.MCPServerEntity) *mcpserver.GetMCPServerListRespData {
	if entity == nil {
		return nil
	}
	var creatorNickname string
	if entity.Creator != nil {
		creatorNickname = entity.Creator.Nickname
	}
	return &mcpserver.GetMCPServerListRespData{
		ID:              entity.ID,
		Name:            entity.Name,
		Description:     entity.Description,
		URL:             entity.Url,
		IsPublic:        entity.IsPublic,
		OpenProxy:       entity.OpenProxy,
		CreatedAt:       entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       entity.UpdatedAt.Format(time.RFC3339),
		CreatorId:       entity.CreatorID,
		CreatorNickname: creatorNickname,
	}
}

func MCPServerEntityToRespData(entity *domain.MCPServerEntity) *mcpserver.GetMCPServerRespData {
	if entity == nil {
		return nil
	}
	tokens := make([]*mcpserver.TokenData, 0, len(entity.Tokens))
	for _, t := range entity.Tokens {
		tokens = append(tokens, &mcpserver.TokenData{
			ID:          t.ID,
			Token:       t.Token,
			Description: t.Description,
		})
	}

	var creatorNickname string
	if entity.Creator != nil {
		creatorNickname = entity.Creator.Nickname
	}
	return &mcpserver.GetMCPServerRespData{
		ID:              entity.ID,
		Name:            entity.Name,
		Description:     entity.Description,
		URL:             entity.Url,
		IsPublic:        entity.IsPublic,
		OpenProxy:       entity.OpenProxy,
		Token:           tokens,
		CreatedAt:       entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       entity.UpdatedAt.Format(time.RFC3339),
		CreatorId:       entity.CreatorID,
		CreatorNickname: creatorNickname,
	}
}

func MCPServerTokenEntityToRespData(entity *domain.MCPServerTokenEntity) *mcpserver.GenerateTokenRespData {
	if entity == nil {
		return nil
	}
	return &mcpserver.GenerateTokenRespData{
		Token: entity.Token,
	}
}
