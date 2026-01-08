# 完整请求流程说明

## 项目架构概览

本项目采用分层架构设计，遵循以下层级结构：
- **Router**: 路由注册和中间件配置
- **Middleware**: 请求拦截和前置处理（如认证、用户信息提取）
- **Handler**: HTTP 请求处理器，负责参数解析和响应返回
- **Service**: 业务逻辑处理层
- **Repo**: 数据持久化层（数据库操作）

## User 请求流程详解

### 流程图
```
客户端请求 (HTTP/RPC)
    ↓
Router (路由匹配)
    ↓
Middleware (认证、权限校验等)
    ↓
Handler (处理器)
    ├─ 解析 Req (HTTP/RPC 参数)
    ├─ 参数验证
    └─ 调用 Service
    ↓
Service (业务逻辑)
    ├─ 转换 Req → Domain Entity
    ├─ 调用 Repo（可能多次）
    ├─ 业务逻辑处理
    └─ 转换 Domain Entity → Resp
    ↓
Repo (数据层)
    ├─ 转换 Domain Entity → DO
    ├─ 执行数据库操作
    └─ 转换 DO → Domain Entity
    ↓
Handler (返回响应)
    └─ 组装 Resp 返回客户端
```

### 具体实现步骤

#### 1. **Router 层** - `biz/router/user/user.go`
```go
// 路由注册
_users.GET("/me", append(_getselfinfoMw(), user.GetSelfInfo)...)
_users.GET("/user-info", append(_getuserinfoMw(), user.GetUserInfo)...)
```

**职责**:
- 定义 HTTP 路由路径和请求方法
- 注册中间件堆栈

#### 2. **Middleware 层** - `biz/router/user/middleware.go`
```go
func _getselfinfoMw() []app.HandlerFunc {
    App := container.GetGlobalContainer()
    return []app.HandlerFunc{
        App.AuthAndUserInfoMiddleware,
    }
}
```

**职责**:
- 认证检查（验证用户是否登录）
- 提取和注入用户上下文信息
- 权限校验

#### 3. **Handler 层** - `biz/handler/user/user_controller.go`

**GetUserInfo 处理流程**:
```go
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
    // 步骤 1: 参数绑定和验证
    var req user.GetUserInfoReq
    err = c.BindAndValidate(&req)
    
    // 步骤 2: 获取全局容器中的 Service
    App := container.GetGlobalContainer()
    
    // 步骤 3: 调用 Service 处理业务逻辑
    resp := App.UserService.GetUserInfo(ctx, &req)
    
    // 步骤 4: 返回响应
    c.JSON(consts.StatusOK, resp)
}
```

**GetSelfInfo 处理流程**:
```go
func GetSelfInfo(ctx context.Context, c *app.RequestContext) {
    // 步骤 1: 参数绑定（GetSelfInfo 无请求参数，用户ID由 Middleware 注入到 ctx）
    var req user.GetUserInfoReq
    err = c.BindAndValidate(&req)
    
    // 步骤 2-4: 同上
}
```

**重要约定**:
- ✅ **只能使用 `model` 包中的 Req/Resp 结构**（这些由 thrift IDL 自动生成）
- ✅ 参数验证和错误返回
- ✅ 调用 Service 进行业务处理
- ❌ 不能直接调用 Repo（通过 Service 代理）
- ❌ 不能处理复杂业务逻辑

#### 4. **Service 层** - `biz/internal/service/user/user.go`

```go
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp) {
    // 步骤 1: 根据请求的用户ID，从 Repo 查询用户信息
    findUser, err := s.Repo.FindUserById(req.UserId)
    
    // 步骤 2: 错误处理
    if err != nil {
        hlog.Errorf("查询用户信息错误:%s", err.Error())
        return &user.GetUserInfoResp{
            Code:    constant.SystemError,
            Message: "系统错误，查询用户信息失败",
        }
    }
    
    // 步骤 3: 业务逻辑处理（可根据需要进行数据转换、计算等）
    
    // 步骤 4: 转换 Domain Entity → Resp
    resp = &user.GetUserInfoResp{
        Code: constant.Success,
        Data: EntityToUserInfoData(findUser),  // 调用 assembler 转换
    }
    return
}
```

**重要约定**:
- ✅ 接收 Handler 传来的 Req（`*user.GetUserInfoReq`）
- ✅ 调用 Repo 获取 Domain Entity（`*domain.UserEntity`）
- ✅ 使用 Domain Entity 进行业务逻辑处理
- ✅ 通过 Assembler 将 Domain Entity 转换为 Resp 中的 Data
- ✅ 返回完整的 Resp（`*user.GetUserInfoResp`）
- ❌ Service 内部只使用 Domain Entity，不直接接触 DO

#### 5. **Assembler 层** - `biz/internal/service/user/assembler.go`

```go
func EntityToUserInfoData(entity *domain.UserEntity) *user.GetUserInfoRespData {
    return &user.GetUserInfoRespData{
        Username: entity.Username,
        Nickname: entity.Nickname,
        Role:     entity.Role,
    }
}
```

**职责**:
- 转换 Domain Entity → Resp Data
- 进行字段映射和数据转换

#### 6. **Repo 层** - `biz/internal/repo/user/user.go`

```go
func (Repo *UserRepoImpl) FindUserById(userId int64) (*domain.UserEntity, error) {
    // 步骤 1: 从数据库查询 DO
    var user do.UserDO
    err := Repo.DB.Where("id = ?", userId).Limit(1).Find(&user).Error
    
    // 步骤 2: 转换 DO → Domain Entity
    return DoToEntity(&user), nil
}

func (Repo *UserRepoImpl) SaveUser(userEntity *domain.UserEntity) error {
    // 步骤 1: 转换 Domain Entity → DO
    userDO := EntityToDo(userEntity)
    
    // 步骤 2: 执行数据库操作
    if userDO.ID == 0 {
        return Repo.DB.Create(&userDO).Error
    } else {
        return Repo.DB.Save(&userDO).Error
    }
}
```

**Assembler 层** - `biz/internal/repo/user/assembler.go`

```go
// DO ↔ Domain Entity 转换
func DoToEntity(do *do.UserDO) *domain.UserEntity { ... }
func EntityToDo(entity *domain.UserEntity) *do.UserDO { ... }
```

**重要约定**:
- ✅ 接收/返回 Domain Entity（`*domain.UserEntity`）
- ✅ 数据库操作时转换为 DO（`*do.UserDO`）
- ✅ 使用 Assembler 进行转换
- ❌ 不返回 DO
- ❌ 不返回 Req/Resp

---

## 数据结构说明

### IDL 定义 - `idl/user.thrift`
```thrift
struct GetUserInfoReq {
    1: required i64 userId (api.query="user_id");
}

struct GetUserInfoRespData {
    1: string username;
    2: string nickname;
    3: string role;
}

struct GetUserInfoResp {
    1: required i32 code;
    2: required string message;
    3: optional GetUserInfoRespData data;
}

service UserController {
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq request) (api.get="/api/users/user-info");
    GetUserInfoResp GetSelfInfo() (api.get="/api/users/me");
}
```

### 数据结构用途

| 位置 | 结构体 | 来源 | 用途 |
|------|--------|------|------|
| Handler | `*user.GetUserInfoReq` | HTTP 请求参数 | 接收客户端请求参数 |
| Handler | `*user.GetUserInfoResp` | Service 返回 | 返回给客户端 |
| Service | `*domain.UserEntity` | Repo 返回 | 业务逻辑处理 |
| Service | `*user.GetUserInfoRespData` | Assembler 转换 | Resp 的 Data 字段 |
| Repo | `*do.UserDO` | 数据库查询结果 | 数据库映射对象 |

### 各层结构体关系

```
Handler 层:
  输入: *user.GetUserInfoReq (来自 HTTP 请求)
  输出: *user.GetUserInfoResp (返回 HTTP 响应)

Service 层:
  输入: *user.GetUserInfoReq
  内部使用: *domain.UserEntity (来自 Repo)
  输出: *user.GetUserInfoResp (通过 Assembler 转换)

Repo 层:
  数据库映射: *do.UserDO
  输出: *domain.UserEntity (通过 Assembler 转换)
```

---

## 新功能开发流程

当需要添加新功能时，按以下步骤进行：

### 1. 修改 IDL 文件
在 `idl/user.thrift` 中添加新的请求/响应结构和服务方法

### 2. 生成代码
运行代码生成工具，自动生成 `biz/model/user/` 中的 Go 代码

### 3. 创建 Router
在 `biz/router/user/user.go` 中注册新的路由，自动生成的代码

### 4. 创建 Middleware
在 `biz/router/user/middleware.go` 中定义新端点的中间件

### 5. 创建 Handler
在 `biz/handler/user/user_controller.go` 中实现新的 HTTP 处理函数：
```go
func NewFeature(ctx context.Context, c *app.RequestContext) {
    var req user.NewFeatureReq
    err := c.BindAndValidate(&req)
    if err != nil {
        // 错误处理
        return
    }
    
    App := container.GetGlobalContainer()
    resp := App.UserService.NewFeature(ctx, &req)
    c.JSON(consts.StatusOK, resp)
}
```

### 6. 创建 Service Interface
在 `biz/internal/service/user/interface.go` 中添加方法：
```go
type IUserService interface {
    NewFeature(ctx context.Context, req *user.NewFeatureReq) (resp *user.NewFeatureResp)
}
```

### 7. 实现 Service
在 `biz/internal/service/user/user.go` 中实现业务逻辑：
```go
func (s *UserServiceImpl) NewFeature(ctx context.Context, req *user.NewFeatureReq) (resp *user.NewFeatureResp) {
    // 调用 Repo 获取数据
    entity, err := s.Repo.FindXXX(...)
    
    // 业务逻辑处理
    
    // 转换并返回
    return &user.NewFeatureResp{...}
}
```

### 8. 创建 Repo Interface（如需要）
在 `biz/internal/repo/user/interface.go` 中添加新的数据访问方法

### 9. 实现 Repo
在 `biz/internal/repo/user/user.go` 中实现数据库操作

### 10. 创建 Domain Entity（如需要）
在 `biz/internal/domain/user.go` 中定义新的实体

### 11. 创建 DO（如需要）
在 `biz/internal/do/user.go` 中定义数据库映射对象

### 12. 创建 Assembler（如需要）
在 `biz/internal/service/user/assembler.go` 和 `biz/internal/repo/user/assembler.go` 中添加转换函数

---

## 关键设计原则

### 分层职责清晰
- **Handler**: HTTP 请求/响应处理，参数验证
- **Service**: 业务逻辑，协调多个 Repo 调用
- **Repo**: 数据持久化，DB 操作
- **Assembler**: 对象转换

### 数据流向明确
```
Request (Req) → Handler → Service → Repo → DO → Database
Response (Resp) ← Handler ← Service ← Repo → Entity
```

### 数据结构隔离
- `Req/Resp`: 只在 Handler 层使用，来自/返回 HTTP
- `Domain Entity`: Service 和 Repo 之间的通讯协议
- `DO`: 仅在 Repo 层使用，对应数据库表结构

### 依赖注入
- 所有依赖通过 Container 注入
- Service 通过 Constructor 接收 Repo 实例
- Middleware 和其他组件通过 Container 获取

---

## 常见错误和注意事项

### ❌ 错误做法
1. Handler 直接调用 Repo
2. Service 返回 DO 或其他数据库类型
3. Repo 返回 Req/Resp
4. 在 Handler 中处理复杂业务逻辑
5. 跨越层级的依赖关系
6. Service 使用多个来源的数据混合返回

### ✅ 正确做法
1. Handler 只调用 Service
2. Service 接收 Req，返回 Resp
3. Service 内部使用 Domain Entity
4. Repo 接收 Domain Entity，返回 Domain Entity
5. 通过 Assembler 进行数据转换
6. 层级之间明确的输入输出契约

---

## 总结

这个架构确保了：
1. **代码复用性**: 同一个 Service 可被多个 Handler 调用
2. **可测试性**: 各层可独立测试
3. **可维护性**: 职责清晰，易于定位问题
4. **可扩展性**: 添加新功能只需沿着流程增加对应层的代码
5. **数据安全**: DO 不会泄露到业务层，隔离数据库实现细节
