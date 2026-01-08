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

service AuthController {
    LoginResp Login(1:LoginReq request) (api.post="/api/auth/login");
}