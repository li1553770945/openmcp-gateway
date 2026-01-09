# QuickStart | 快速开始

本指南将指引你在几分钟内完成 **OpenMCP-Gateway** 后端的部署与运行。

## 🛠 环境要求

在开始之前，请确保你的环境满足以下条件：

* **操作系统**: Linux 或 macOS (强烈建议 **不要** 使用 Windows)。
* **Go 语言**: [Go 1.25.1](https://gs.jurieo.com/gemini/official/search?q=https://go.dev/dl/) 及以上版本。
* **网络环境**: 确保服务器防火墙已放行你准备配置的 API 端口。

为了方便没有 Go 语言环境的用户，我们提供[release版本](https://github.com/li1553770945/openmcp-gateway/releases)，但不保证release的更新速度，建议使用源码编译版本。

---

## 🚀 部署步骤

如果你使用release，请跳过编译步骤，直接下载对应平台的二进制文件，解压后进行配置和启动。

### 1. 克隆代码库

首先，将项目源码克隆到本地或服务器：

```bash
git clone https://github.com/li1553770945/openmcp-gateway
cd openmcp-gateway

```

### 2. 初始化环境

下载项目所需的依赖包：

```bash
go mod tidy

```

### 3. 配置文件准备

项目通过 `conf/` 目录下的配置文件管理运行参数。请根据你的运行环境进行设置：

1. **创建配置目录**: `mkdir -p conf`
2. **复制示例配置**: `cp config-example.yml conf/`
3. **重命名配置文件**:
* **开发环境**: 重命名为 `conf/deployment.yml`
* **生产环境**: 重命名为 `conf/production.yml`



> 📌 **注意**: 具体的配置项含义（如端口、数据库、密钥等），请参考 [配置说明文档](/docs/configuration.md)。

### 4. 启动服务

启动时必须通过环境变量 `ENV` 指定配置文件。

**开发模式 (Development):**

```bash
export ENV=development
go run .

```

**生产模式 (Production):**

```bash
export ENV=production
go run .

```

---

## ✅ 结果验证

服务启动后，可以通过以下两种方式验证是否部署成功：

### 1. 基础连通性测试

在终端执行：

```bash
curl http://127.0.0.1:<你的端口>/ping

```

**期望响应**: `{"code": 0, "message": "pong"}`

### 2. 交互式文档访问

在浏览器中打开：
`http://<服务器IP>:<端口>/docs`
如果能正常看到 Swagger 或 ReDoc 界面，说明后端服务及文档系统均已就绪。

---

## 💡 常见问题排查

* **报错 `config file not found**`: 请检查 `ENV` 变量是否设置，且 `conf/` 目录下是否存在对应的 `.yml` 文件。

---

## 🔐 接口调用示范 (Usage)

本服务采用 **JWT (JSON Web Token)** 认证机制。在访问受保护的接口前，需要先通过登录接口获取 Token。

### 1. 获取 Token (登录)

调用 `/api/auth/login` 接口获取访问凭证：

```bash
curl -X POST http://127.0.0.1:9000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "your_username",
    "password": "your_password"
  }'
```

**响应示例:**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 2. 访问受保护接口

在后续请求中，需将获取到的 Token 添加到 HTTP Header 的 `Authorization` 字段中，格式为 `Bearer <token>`：

```bash
curl http://127.0.0.1:9000/api/protected/resource \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```


