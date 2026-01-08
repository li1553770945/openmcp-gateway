namespace go user

struct GetUserInfoReq {
    // 用户ID
    1: required i64 userId (api.query="user_id");
}

struct GetUserInfoRespData {
    // 用户名
    1: string username;
    // 昵称
    2: string nickname;
    // 角色
    3: string role;
}

struct GetUserInfoResp{
    // 状态码
    1: required i32 code;
    // 消息
    2: required string message;
    // 数据
    3: optional GetUserInfoRespData data;
}


service UserController {
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq request) (api.get="/api/users/user-info");
    GetUserInfoResp GetSelfInfo() (api.get="/api/users/me");
}