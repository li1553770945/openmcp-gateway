namespace go user

// 注册用户
struct RegisterReq {
    // 用户名
    1: required string username,
    // 密码
    2: required string password,
    // 邮箱
    3: required string email,
    // 昵称
    4: required string nickname,
}

// 注册用户响应
struct RegisterResp {
    // 状态码
    1: required i32 code,
    // 消息
    2: required string message,
}

struct GetUserInfoReq {
    // 用户ID
    1: required i64 userId (api.path = "id"),
}

struct GetUserInfoRespData {
    // 用户ID
    1: required i64 id,
    // 用户名
    2: required string username,
    // 昵称
    3: required string nickname,
    // 角色
    4: required string role,
}

struct GetUserInfoResp {
    // 状态码
    1: required i32 code,
    // 消息
    2: required string message,
    // 数据
    3: optional GetUserInfoRespData data,
}

struct UpdateSelfInfoReq {
    // 昵称
    1: optional string nickname,
    // 密码
    2: optional string password,
}

struct UpdateSelfInfoResp {
    // 状态码
    1: required i32 code,
    // 消息
    2: required string message,
}

service UserController {
    RegisterResp Register(1: RegisterReq request) (api.post = "/api/users"),
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq request) (api.get = "/api/users/:id"),
    GetUserInfoResp GetSelfInfo() (api.get = "/api/users/me"),
    UpdateSelfInfoResp UpdateSelfInfo(1: UpdateSelfInfoReq request) (api.put = "/api/users/me"),
}