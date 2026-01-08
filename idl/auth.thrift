namespace go auth

struct LoginReq {
     1: required string username;
     2: required string password;
}

struct LoginRespData {
    1: string token;
}

struct LoginResp{
    1: required i32 code;
    2: required string message;
    3: optional LoginRespData data;
}

struct RegisterReq {
     1: required string username;
     2: required string password;
     3: required string email;
}
struct RegisterResp {
    1: required i32 code;
    2: required string message;
}

service AuthController {
    LoginResp Login(1:LoginReq request) (api.post="/api/auth/login");
    RegisterResp Register(1:RegisterReq request) (api.post="/api/auth/register");
}