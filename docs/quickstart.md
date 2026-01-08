# QuickStart

本指南将帮助你快速部署和运行 openmcp 端口转发工具-后端。

## 环境准备

1. 一台服务器或本地机器，建议使用服务器
2. linux或 macOS 操作系统（推荐使用linux，已知在windows上很可能存在问题，强烈建议不要使用windows）
3. Go1.25.1 及以上版本

## 部署步骤


1. 克隆代码库

```bash
git clone
cd openmcp-gateway
```

2. 安装依赖

```bash
go mod tidy
```

3. 新建配置文件

新建一个`conf`文件夹，复制 `config-example.yml` 文件到该文件夹并重命名。

本地开发环境重命名为 `deployment.yml`，生产环境重命名为 `production.yml`，这将根据`ENV`环境变量来决定加载哪个配置文件。

有关配置文件的详细信息，请参考[配置说明](./docs/configuration.md)。

4. 运行

```bash
export ENV=development # 或 production 根据需要选择环境，必须设置且必须是二者之一
go run .
```

5. 访问

访问`http://<服务器IP>:<端口>/ping`，你应该能看到输出`{"message":"pong"}`。