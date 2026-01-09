package domain

import "time"

type MCPServerEntity struct {
	ID          int64
	Name        string
	Description string
	Url         string
	IsPublic    bool
	OpenProxy   bool
	CreatorID   int64
	Creator     *UserEntity
	Tokens      []MCPServerTokenEntity
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MCPServerTokenEntity struct {
	ID          int64
	Token       string
	Description string
	MCPServerID int64
}
