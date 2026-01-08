package mcpserver

import (
	"errors"

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
	serverDO := EntityToDo(server)
	if serverDO.ID == 0 {
		return r.DB.Create(serverDO).Error
	} else {
		return r.DB.Omit("CreatedAt", "DeletedAt").Save(serverDO).Error
	}
}

func (r *MCPServerRepoImpl) FindMCPServerById(id int64) (*domain.MCPServerEntity, error) {
	var serverDO do.MCPServerDO
	err := r.DB.Preload("Tokens").Where("id = ?", id).First(&serverDO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return DoToEntity(&serverDO), nil
}

func (r *MCPServerRepoImpl) ListMCPServersByCreatorId(creatorId int64, start, end int64) ([]*domain.MCPServerEntity, error) {
	var serverDOs []do.MCPServerDO
	limit := int(end - start)
	if limit <= 0 {
		limit = 10
	}
	offset := int(start)

	err := r.DB.Where("creator_id = ?", creatorId).
		Offset(offset).Limit(limit).
		Order("id desc").
		Find(&serverDOs).Error

	if err != nil {
		return nil, err
	}

	results := make([]*domain.MCPServerEntity, len(serverDOs))
	for i, s := range serverDOs {
		results[i] = DoToEntity(&s)
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
		Offset(offset).Limit(limit).
		Order("id desc").
		Find(&serverDOs).Error

	if err != nil {
		return nil, err
	}

	results := make([]*domain.MCPServerEntity, len(serverDOs))
	for i, s := range serverDOs {
		results[i] = DoToEntity(&s)
	}
	return results, nil
}

func (r *MCPServerRepoImpl) SaveToken(token *domain.MCPServerTokenEntity) error {
	tokenDO := TokenEntityToDo(token)
	if tokenDO.ID == 0 {
		return r.DB.Create(tokenDO).Error
	}
	return r.DB.Save(tokenDO).Error
}
