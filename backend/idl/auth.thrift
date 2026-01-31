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


service AuthController {
    LoginResp Login(1:LoginReq request) (api.post="/api/auth/login");
}