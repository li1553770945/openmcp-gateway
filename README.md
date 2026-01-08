# openmcp 端口转发工具-后端

## QuickStart

如果你不想参与开发，只是想要部署和使用本项目，请直接参考[QuickStart](./docs/QuickStart.md)

## 开发指南

### 根据idl更新接口

```bash

hz update -idl idl/user.thrift   
hz update -idl idl/auth.thrift   
hz update -idl idl/mcpserver.thrift
```