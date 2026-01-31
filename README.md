# OpenMCP Gateway

<p align="center">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License">
  <img src="https://img.shields.io/badge/Status-Active-brightgreen.svg" alt="Status">
</p>

**OpenMCP Gateway** 是一个前后端一体的开源项目，用于解决复杂网络环境下的 MCP Server 暴露问题，提供 RESTful API、管理控制台、鉴权与代理能力。

本仓库为**单体仓库（Monorepo）**，包含后端服务与前端管理控制台。

---

## 📁 项目结构

```
openmcp-gateway/
├── backend/          # 后端服务（Go）
│   ├── docs/         # 后端文档（快速开始、配置、开发指南等）
│   ├── idl/          # Thrift 接口定义
│   └── ...
├── frontend/         # 前端管理控制台（Next.js + React）
│   ├── app/          # 页面与路由
│   ├── components/   # UI 组件
│   └── ...
├── .github/          # CI/CD（当前为后端构建与 Release）
│   └── workflows/
│       └── build.yaml
└── README.md         # 本文件
```

- **`backend/`**：Go 后端，负责 API、鉴权、MCP 代理等，详见 [backend/README.md](./backend/README.md)。
- **`frontend/`**：Next.js 前端，用于管理 MCP 服务器、用户与令牌等，详见 [frontend/README.md](./frontend/README.md)。
- **`.github/`**：来自原后端仓库，目前工作流仅构建并发布**后端**二进制（推送到 `master` 或打 `v*` tag 时触发）。

---

## 🚀 快速开始

### 1. 克隆仓库

```bash
git clone https://github.com/li1553770945/openmcp-gateway.git
cd openmcp-gateway
```

### 2. 启动后端

```bash
cd backend
go mod tidy
# 按 backend/docs/quickstart.md 准备 conf 与 config（如复制 config-example.yml）
export ENV=development
go run .
```

默认 API 地址：`http://localhost:9000`，Swagger 文档：`http://localhost:9000/docs`。

### 3. 启动前端（可选）

```bash
cd frontend
yarn install
# 新建 .env.local，设置 NEXT_PUBLIC_API_BASE_URL=http://localhost:9000/api
yarn dev
```

浏览器访问前端开发地址（通常为 `http://localhost:3000`）。

---

## 📖 更多说明

| 内容           | 说明 |
|----------------|------|
| 后端部署与配置 | [backend/docs/quickstart.md](./backend/docs/quickstart.md)、[backend/docs/configuration.md](./backend/docs/configuration.md) |
| 后端开发指南   | [backend/docs/development.md](./backend/docs/development.md) |
| 前端技术栈与运行 | [frontend/README.md](./frontend/README.md) |
| API 响应格式与状态码 | [backend/README.md#API 文档规范](./backend/README.md) |
| JWT 认证用法   | [backend/README.md#认证规范](./backend/README.md) |

---

## 🔧 技术栈概览

- **后端**：Go（Hertz 等）、SQLite/数据库、JWT、Thrift IDL。
- **前端**：Next.js（App Router）、TypeScript、Tailwind CSS、Shadcn UI、Zustand、SWR。

---

## 📜 开源协议

本项目基于 [MIT License](https://opensource.org/licenses/MIT) 开源。

---

## 🤝 参与贡献

欢迎提交 Issue 与 Pull Request。开发前请阅读 [backend/docs/development.md](./backend/docs/development.md)。

<p align="center">如果这个项目对你有帮助，欢迎给一个 ⭐️ Star。</p>
