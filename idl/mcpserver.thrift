namespace go mcpserver

struct AddMCPServerReq {
    1: required string username;
    2: required string password;
}

struct AddMCPServerRespData {
    1: required string name;
    2: required string url;
}

struct AddMCPServerResp{
     1: required i32 code;
     2: required string message;
     3: optional AddMCPServerRespData data;

}


service MCPServerService {
    AddMCPServerResp AddMCPServer(1: AddMCPServerReq req)(api.post="/api/mcpserver");
}