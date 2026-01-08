package do

type MCPServerDO struct {
	BaseModel
	Name        string             `gorm:"type:varchar(255)"`
	Description string             `gorm:"type:text"`
	Url         string             `gorm:"type:varchar(255)"`
	IsPublic    bool               `gorm:"default:false"`
	OpenProxy   bool               `gorm:"default:false"`
	CreatorID   int64              `gorm:"index"`
	Tokens      []MCPServerTokenDO `gorm:"foreignKey:MCPServerID"`
}

type MCPServerTokenDO struct {
	BaseModel
	Token       string `gorm:"type:varchar(255);uniqueIndex"`
	Description string `gorm:"type:text"`
	MCPServerID int64  `gorm:"index"`
}
