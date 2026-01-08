namespace go auth

struct LoginReq {
     // 用户名
     1: required string username;
     // 密码
     2: required string password;
}

struct LoginRespData {
    // 登录凭证
    1: string token;
}

struct LoginResp{
    // 状态码
    1: required i32 code;
    // 消息
    2: required string message;
    // 数据
    3: optional LoginRespData data;
}

struct RegisterReq {
     // 用户名
     1: required string username;
     // 密码
     2: required string password;
     // 邮箱
     3: required string email;
}
struct RegisterResp {
    // 状态码
    1: required i32 code;
    // 消息
    2: required string message;
}

service AuthController {
    LoginResp Login(1:LoginReq request) (api.post="/api/auth/login");
    RegisterResp Register(1:RegisterReq request) (api.post="/api/auth/register");
}