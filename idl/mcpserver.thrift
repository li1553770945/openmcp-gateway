namespace go mcpserver

// 增加MCPServer
struct AddMCPServerReq {
    // MCPServer 名称
    1: required string name;
    // MCPServer 描述
    2: required string description;
    // MCPServer URL
    3: required string url;
    // 是否公开
    4: required bool isPublic;
    // 是否开启代理
    5: required bool openProxy;
}
struct AddMCPServerResp{
     // 状态码
     1: required i32 code;
     // 消息
     2: required string message;
}

// 生成token
struct GenerateTokenReq{
    // MCPServer ID
    1: required i64 id;
    // Token 描述
    2: required string description;
}

struct GenerateTokenRespData{
    // 生成的 Token
    1: required string token
}

struct GenerateTokenResp{
    // 状态码
    1: required i32 code;
    // 消息
    2: required string message;
    // 数据
    3: optional GenerateTokenRespData data;
}

// 获取单个MCPServer

struct GetMCPServerReq{
    // MCPServer ID
    1: required i64 mcpServerId (api.path="id");
}
struct TokenData{
    // Token
    1: required string token;
    // 描述
    2: required string description;
}

struct GetMCPServerRespData{
    // MCPServer ID
    1: required i64 id;
    // 名称
    2: required string name;
    // 描述
    3: required string description;
    // URL
    4: required string url;
    // 是否公开
    5: required bool isPublic;
    // 是否开启代理
    6: required bool openProxy;
    // Token列表
    7: required list<TokenData> token;
    // 创建时间
    8: required string createdAt;
    // 更新时间
    9: required string updatedAt;
}

struct GetMCPServerResp{
    // 状态码
    1: required i32 code;
    // 消息
    2: required string message;
    // 数据
    3: optional GetMCPServerRespData data;
}


// 获取MCPServer列表

struct GetMCPServerListReq{
    // 起始位置
    1: required i64 start;
    // 结束位置
    2: required i64 end;
}

struct GetMCPServerListRespData{
    // MCPServer ID
    1: required i64 id;
    // 名称
    2: required string name;
    // 描述
    3: required string description;
    // URL
    4: required string url;
    // 是否公开
    5: required bool isPublic;
    // 是否开启代理
    6: required bool openProxy;
    // 创建时间
    7: required string createdAt;
    // 更新时间
    8: required string updatedAt;
}


struct GetMCPServerListResp{
    // 状态码
    1: required i32 code;
    // 消息
    2: required string message;
    // 数据列表
    3: list<GetMCPServerListRespData> data;
}


// 更新现有MCPServer
struct UpdateMCPServerReq {
    // MCPServer ID
    1: required i64 id (api.path="id");
    // 名称
    2: required string name;
    // 描述
    3: required string description;
    // URL地址
    4: required string url;
    // 是否公开
    5: required bool isPublic;
    // 是否开启代理
    6: required bool openProxy;
}

struct UpdateMCPServerResp{
     // 状态码
     1: required i32 code;
     // 消息
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