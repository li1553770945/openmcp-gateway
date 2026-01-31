# OpenMCP Gateway Frontend

OpenMCP Gateway 的前端管理控制台。基于 [Next.js](https://nextjs.org/) 和 [React](https://react.dev/) 构建，用于管理 MCP 服务器、生成访问令牌以及用户管理。

OpenMCP Gateway 后端仓库: [openmcp-gateway](https://github.com/li1553770945/openmcp-gateway)

## 🛠 技术栈

- **框架**: [Next.js 22+](https://nextjs.org/) (App Router)
- **语言**: TypeScript
- **样式**: [Tailwind CSS](https://tailwindcss.com/)
- **组件库**: [Shadcn UI](https://ui.shadcn.com/) 或其他基于 Shadcn UI 的组件库
- **状态管理**: Zustand
- **数据请求**: SWR 

## 🚀 快速开始

### 1. 环境准备

确保你的本地环境已安装：
- Node.js >= 22.18.0
- yarn 

### 2. 安装依赖

```bash
yarn install
```

### 3. 配置环境变量

创建 `.env.local` 文件并添加以下内容：

```env
NEXT_PUBLIC_API_BASE_URL=http://localhost:9000/api
```

根据你的后端 API 地址修改 `NEXT_PUBLIC_API_BASE_URL`。

### 4. 运行开发服务器

```bash
yarn dev
```