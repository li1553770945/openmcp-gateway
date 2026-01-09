package converter

import (
	"github.com/li1553770945/openmcp-gateway/biz/internal/do"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
)

func MCPServerDoToEntity(serverDO *do.MCPServerDO) *domain.MCPServerEntity {
	if serverDO == nil {
		return nil
	}
	tokens := make([]domain.MCPServerTokenEntity, 0, len(serverDO.Tokens))
	for _, tokenDO := range serverDO.Tokens {
		tokens = append(tokens, *MCPServerTokenDoToEntity(&tokenDO))
	}

	return &domain.MCPServerEntity{
		ID:          serverDO.ID,
		Name:        serverDO.Name,
		Description: serverDO.Description,
		Url:         serverDO.Url,
		IsPublic:    serverDO.IsPublic,
		OpenProxy:   serverDO.OpenProxy,
		CreatorID:   serverDO.CreatorID,
		Tokens:      tokens,
		CreatedAt:   serverDO.CreatedAt,
		UpdatedAt:   serverDO.UpdatedAt,
		Creator:     UserDoToEntity(serverDO.Creator),
	}
}

func MCPServerEntityToDo(serverEntity *domain.MCPServerEntity) *do.MCPServerDO {
	if serverEntity == nil {
		return nil
	}
	tokens := make([]do.MCPServerTokenDO, 0, len(serverEntity.Tokens))
	for _, tokenEntity := range serverEntity.Tokens {
		tokens = append(tokens, *MCPServerTokenEntityToDo(&tokenEntity))
	}

	return &do.MCPServerDO{
		BaseModel: do.BaseModel{
			ID:        serverEntity.ID,
			CreatedAt: serverEntity.CreatedAt,
			UpdatedAt: serverEntity.UpdatedAt,
		},
		Name:        serverEntity.Name,
		Description: serverEntity.Description,
		Url:         serverEntity.Url,
		IsPublic:    serverEntity.IsPublic,
		OpenProxy:   serverEntity.OpenProxy,
		CreatorID:   serverEntity.CreatorID,
		Creator:     UserEntityToDo(serverEntity.Creator),
		Tokens:      tokens,
	}
}

func MCPServerTokenDoToEntity(tokenDO *do.MCPServerTokenDO) *domain.MCPServerTokenEntity {
	if tokenDO == nil {
		return nil
	}
	return &domain.MCPServerTokenEntity{
		ID:          tokenDO.ID,
		Token:       tokenDO.Token,
		Description: tokenDO.Description,
		MCPServerID: tokenDO.MCPServerID,
	}
}

func MCPServerTokenEntityToDo(tokenEntity *domain.MCPServerTokenEntity) *do.MCPServerTokenDO {
	if tokenEntity == nil {
		return nil
	}
	return &do.MCPServerTokenDO{
		BaseModel: do.BaseModel{
			ID: tokenEntity.ID,
		},
		Token:       tokenEntity.Token,
		Description: tokenEntity.Description,
		MCPServerID: tokenEntity.MCPServerID,
	}
}
