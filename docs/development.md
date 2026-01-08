# 开发指南 (Development Guide)

感谢你对 **OpenMCP-Gateway** 的关注！本指南将帮助你快速搭建开发环境并了解项目的工程化流程。

## 🛠 开发环境准备

在开始贡献代码前，请确保你的本地环境已安装以下工具：

* **Go**: 1.25.1+
* **Hertz 命令行工具 (hz)**: [安装指南](https://www.cloudwego.io/zh/docs/hertz/getting-started/)
```bash
go install github.com/cloudwego/hertz/cmd/hz@latest

```


* **Swag**: 用于生成接口文档
```bash
go install github.com/swaggo/swag/cmd/swag@latest

```



---

## 🏗 核心开发流程

本项目采用 **IDL (Interface Definition Language)** 驱动开发，所有接口定义均维护在 `idl/` 目录下。

### 1. 更新接口定义 (Thrift)

当你修改了 `idl/*.thrift` 文件后，需要使用 `hz` 工具同步生成底层的桩代码（Stub）和 Handler 模板：

```bash
# 更新用户模块接口
hz update -idl idl/user.thrift   

# 更新认证模块接口
hz update -idl idl/auth.thrift   

# 更新 MCP Server 模块接口
hz update -idl idl/mcpserver.thrift

```

> **寓意标注：**
> * **IDL 目录**：白色直线指向 `idl/` —— **“契约之源”**，定义系统间交流的标准语言。
> * **hz update**：白色直线指向生成命令 —— **“自动化之手”**，将逻辑抽象转化为高效代码。
> 
> 

### 2. 同步 API 文档

为了保证 `http://localhost:port/docs` 显示的内容与代码一致，在修改 Handler 注释后，请务必更新 Swagger 文档：

```bash
swag init -o docs/swagger

```

---

## 📂 项目目录结构说明

```text
.
├── biz/                  # 业务逻辑根目录
│   ├── constant/         # 全局常量 (错误码、环境变量名)
│   ├── container/        # 全局依赖容器 (IoC)
│   ├── handler/          # 请求处理函数 (Controller, 处理 HTTP 请求，由 hz 生成)
│   ├── infra/            # 基础设施支持
│   │   ├── config/       # 配置管理
│   │   └── database/     # 数据库初始化
│   ├── internal/         # 核心业务逻辑 (DDD 风格)
│   │   ├── converter/    # 数据转换 (Assembler)
│   │   ├── do/           # 数据库实体 (Data Objects)
│   │   ├── domain/       # 领域实体 (Domain Model)
│   │   ├── middleware/   # 中间件逻辑
│   │   ├── repo/         # 仓储层实现 (Repository)
│   │   └── service/      # 业务服务实现 (Service)
│   ├── model/            # IDL 生成的数据模型 (DTO)
│   └── router/           # 路由注册
├── conf/                 # 配置文件目录
├── docs/                 # 项目文档及 Swagger 资源
├── idl/                  # Thrift 接口定义文件 (Source of Truth)
├── output/               # 构建产出目录
├── script/               # 启动与辅助脚本
├── go.mod                # 依赖管理
└── main.go               # 程序入口
```

---

## 🤝 代码提交规范

我们建议在提交 Pull Request 前遵循以下准则：

1. **格式化**: 执行 `go fmt ./...` 保持代码整洁。
2. **注释**: 复杂的业务逻辑请务必添加中文注释。
3. **测试**: 建议在 `biz/service` 层添加必要的单元测试。
