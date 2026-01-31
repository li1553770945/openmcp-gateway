# <p align="center">OpenMCP-Gateway Backend</p>

<p align="center">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License">
  <img src="https://img.shields.io/badge/Status-Active-brightgreen.svg" alt="Status">
</p>

**OpenMCP-Gateway** 是一个高性能、易配置的开源转发工具。它致力于解决复杂网络环境下的MCPServer暴露问题，为开发者提供简洁、稳定且安全的流量转发方案。

前端项目地址：👉 **[OpenMCP-Gateway Frontend](https://github.com/li1553770945/openmcp-gateway-fe)**。

---

## 📖 项目简介

OpenMCP-Gateway Backend 负责处理核心转发逻辑以及API 管理。采用单体轻量化架构设计，旨在通过极简的配置实现复杂的转发需求。

### 核心特性
<!-- * 🚀 **高性能转发**：优化底层连接池，实现低延迟、高吞吐的流量处理。 -->
* 🛠️ **标准 RESTful API**：提供规范的接口，方便集成到自动化运维系统或自定义 UI。
* 🔐 **安全防护**：内置鉴权机制，支持自定义访问策略，确保数据传输安全。
<!-- * 📊 **可视化监控**：支持实时连接数监控与流量统计。 -->

---

## 🚀 快速开始

### 环境准备
* **操作系统**: Linux, macOS（Windows未经测试，无法保证稳定运行）
* **基础依赖**: Go 1.25.1 及以上版本

### 部署使用

如果你不需要参与代码开发，仅需部署服务，请直接参考：
👉 **[快速开始部署手册](./docs/quickstart.md)**

---

## 📋 API 文档规范

本项目所有接口响应均遵循统一的 JSON 格式标准。

### 通用响应格式
```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

根据不同的请求结果，data 字段会包含相应的响应数据，具体请参考各接口说明。


### 状态码定义

| 状态码 (code) | 语义       | 详细描述                         |
| :------------ | :--------- | :------------------------------- |
| `0`           | 成功       | 请求已成功处理                   |
| `4001`        | 参数错误   | 客户端请求参数不合法或缺失       |
| `4003`        | 未授权     | 请求未携带有效的身份认证信息     |
| `4004`        | 资源不存在 | 请求的 API 路径或特定资源未找到  |
| `5001`        | 系统错误   | 服务端内部处理异常               |


### 在线交互文档

服务部署后，请访问 `http://<server_ip>:<port>/docs` 查看完整的 Swagger/ReDoc 接口说明。



---

## 🔐 认证规范 (Usage)

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


## 🛠️ 开发指南

我们非常欢迎开发者提交 Pull Request 或报告 Issue！

1. Fork 本仓库。

2. 详细阅读 [开发搭建指南](./docs/development.md)。

3. 创建你的特性分支 (git checkout -b feature/AmazingFeature)。

4. 提交你的修改 (git commit -m 'feat: add some amazing feature')。

5. 推送到分支 (git push origin feature/AmazingFeature)。  

6. 发起 Pull Request。

## 📜 开源协议

本项目基于 MIT License 协议开源。

## 🤝 联系与支持

报告 Bug 或建议: [GitHub Issues](https://github.com/li1553770945/openmcp-gateway/issues)

贡献代码: 欢迎任何形式的 Pull Request！

<p align="center">如果这个项目对你有帮助，请给一个 ⭐️ Star 吧！</p>