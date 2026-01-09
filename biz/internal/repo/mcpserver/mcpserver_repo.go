package mcpserver

import (
	"errors"

	assembler "github.com/li1553770945/openmcp-gateway/biz/internal/converter"
	"github.com/li1553770945/openmcp-gateway/biz/internal/do"
	"github.com/li1553770945/openmcp-gateway/biz/internal/domain"
	"gorm.io/gorm"
)

type MCPServerRepoImpl struct {
	DB *gorm.DB
}

func NewMCPServerRepository(db *gorm.DB) IMCPServerRepository {
	err := db.AutoMigrate(&do.MCPServerDO{}, &do.MCPServerTokenDO{})
	if err != nil {
		panic("迁移MCPServer模型失败：" + err.Error())
	}
	return &MCPServerRepoImpl{
		DB: db,
	}
}

func (r *MCPServerRepoImpl) SaveMCPServer(server *domain.MCPServerEntity) error {
	serverDO := assembler.MCPServerEntityToDo(server)
	if serverDO.ID == 0 {
		return r.DB.Create(serverDO).Error
	} else {
		return r.DB.Omit("CreatedAt", "DeletedAt").Save(serverDO).Error
	}
}

func (r *MCPServerRepoImpl) FindMCPServerById(id int64) (*domain.MCPServerEntity, error) {
	var serverDO do.MCPServerDO
	err := r.DB.Preload("Tokens").Preload("Creator").Where("id = ?", id).First(&serverDO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return assembler.MCPServerDoToEntity(&serverDO), nil
}

func (r *MCPServerRepoImpl) ListMCPServersByCreatorId(creatorId int64, start, end int64) ([]*domain.MCPServerEntity, error) {
	var serverDOs []do.MCPServerDO
	limit := int(end - start)
	if limit <= 0 {
		limit = 10
	}
	offset := int(start)

	err := r.DB.Where("creator_id = ?", creatorId).
		Preload("Creator").
		Offset(offset).Limit(limit).
		Order("id desc").
		Find(&serverDOs).Error

	if err != nil {
		return nil, err
	}

	results := make([]*domain.MCPServerEntity, len(serverDOs))
	for i, s := range serverDOs {
		results[i] = assembler.MCPServerDoToEntity(&s)
	}
	return results, nil
}

func (r *MCPServerRepoImpl) ListPublicMCPServers(start, end int64) ([]*domain.MCPServerEntity, error) {
	var serverDOs []do.MCPServerDO
	limit := int(end - start)
	if limit <= 0 {
		limit = 10
	}
	offset := int(start)

	err := r.DB.Where("is_public = ?", true).
		Preload("Creator").
		Offset(offset).Limit(limit).
		Order("id desc").
		Find(&serverDOs).Error

	if err != nil {
		return nil, err
	}

	results := make([]*domain.MCPServerEntity, len(serverDOs))
	for i, s := range serverDOs {
		results[i] = assembler.MCPServerDoToEntity(&s)
	}
	return results, nil
}

func (r *MCPServerRepoImpl) SaveToken(token *domain.MCPServerTokenEntity) error {
	tokenDO := assembler.MCPServerTokenEntityToDo(token)
	if tokenDO.ID == 0 {
		return r.DB.Create(tokenDO).Error
	}
	return r.DB.Save(tokenDO).Error
}

func (r *MCPServerRepoImpl) FindTokenByToken(token string) (*domain.MCPServerTokenEntity, error) {
	var tokenDO do.MCPServerTokenDO
	// 使用 Find 替代 First 以避免记录不存在时的日志警告
	if err := r.DB.Where("token = ?", token).Limit(1).Find(&tokenDO).Error; err != nil {
		return nil, err
	}
	if tokenDO.ID == 0 {
		return nil, nil
	}
	return assembler.MCPServerTokenDoToEntity(&tokenDO), nil
}

func (r *MCPServerRepoImpl) FindTokenById(id int64) (*domain.MCPServerTokenEntity, error) {
	var tokenDO do.MCPServerTokenDO
	// 使用 Find 替代 First 以避免记录不存在时的日志警告
	if err := r.DB.Where("id = ?", id).Limit(1).Find(&tokenDO).Error; err != nil {
		return nil, err
	}
	if tokenDO.ID == 0 {
		return nil, nil
	}
	return assembler.MCPServerTokenDoToEntity(&tokenDO), nil
}

func (r *MCPServerRepoImpl) DeleteMCPServer(id int64) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 删除相关的 token
		if err := tx.Where("mcp_server_id = ?", id).Delete(&do.MCPServerTokenDO{}).Error; err != nil {
			return err
		}
		// 删除服务器
		if err := tx.Delete(&do.MCPServerDO{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *MCPServerRepoImpl) DeleteToken(id int64) error {
	return r.DB.Delete(&do.MCPServerTokenDO{}, id).Error
}

func (r *MCPServerRepoImpl) FindTokensByMcpServerId(mcpServerId int64) ([]*domain.MCPServerTokenEntity, error) {
	var tokenDOs []do.MCPServerTokenDO
	err := r.DB.Where("mcp_server_id = ?", mcpServerId).Find(&tokenDOs).Error
	if err != nil {
		return nil, err
	}
	results := make([]*domain.MCPServerTokenEntity, len(tokenDOs))
	for i, t := range tokenDOs {
		results[i] = assembler.MCPServerTokenDoToEntity(&t)
	}
	return results, nil
}
