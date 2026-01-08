namespace go mcpserver

// 增加MCPServer
struct AddMCPServerReq {
    1: required string name;
    2: required string description;
    3: required string url;
    4: required bool isPublic;
    5: required bool openProxy;
}
struct AddMCPServerResp{
     1: required i32 code;
     2: required string message;
}

// 生成token
struct GenerateTokenReq{
    1: required i64 id;
    2: required string description;
}

struct GenerateTokenRespData{
    1: required string token
}

struct GenerateTokenResp{
    1: required i32 code;
    2: required string message;
    3: optional GenerateTokenRespData data;
}

// 获取单个MCPServer

struct GetMCPServerReq{
    1: required i64 mcpServerId (api.path="id");
}
struct TokenData{
    1: required string token;
    2: required string description;
}

struct GetMCPServerRespData{
    1: required string name;
    2: required string description;
    3: required string url;
    4: required bool isPublic;
    5: required bool openProxy;
    6: required list<TokenData> token;
}

struct GetMCPServerResp{
    1: required i32 code;
    2: required string message;
    3: optional GetMCPServerRespData data;
}


// 获取MCPServer列表

struct GetMCPServerListReq{
    1: required i64 start;
    2: required i64 end;
}

struct GetMCPServerListRespData{
    1: required string name;
    2: required string description;
    3: required string url;
    4: required bool isPublic;
    5: required bool openProxy;
}


struct GetMCPServerListResp{
    1: required i32 code;
    2: required string message;
    3: optional list<GetMCPServerListRespData> data;
}


// 更新现有MCPServer
struct UpdateMCPServerReq {
    1: required i64 id (api.path="id");
    2: required string name;
    3: required string description;
    4: required string url;
    5: required bool isPublic;
    6: required bool openProxy;
}

struct UpdateMCPServerResp{
     1: required i32 code;
     2: required string message;
}



service MCPServerService {
    AddMCPServerResp AddMCPServer(1: AddMCPServerReq req)(api.post="/api/mcpservers");
    GenerateTokenResp GenerateToken(1: GenerateTokenReq req)(api.post="/api/mcpservers/generate-token")
    GetMCPServerListResp GetSelfMCPServerList(1: GetMCPServerListReq req)(api.get="/api/mcpservers/self")
    GetMCPServerListResp GetPublicMCPServerList(1: GetMCPServerListReq req)(api.get="/api/mcpservers/public")
    UpdateMCPServerResp UpdateMCPServer(1: UpdateMCPServerReq req)(api.put="/api/mcpservers/:id")

    GetMCPServerResp GetMCPServer(1:GetMCPServerReq req)(api.get="/api/mcpservers/:id")
}