package domain

type MCPServerEntity struct {
	ID          int64
	Name        string
	Description string
	Url         string
	IsPublic    bool
	OpenProxy   bool
	CreatorID   int64
	Tokens      []MCPServerTokenEntity
}

type MCPServerTokenEntity struct {
	ID          int64
	Token       string
	Description string
	MCPServerID int64
}
