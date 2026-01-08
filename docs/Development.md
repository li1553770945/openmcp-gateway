# 开发指南
## 根据idl更新接口

如果您修改了`idl`目录下的`.thrift`文件，请运行以下命令更新接口代码：

```bash

hz update -idl idl/user.thrift   
hz update -idl idl/auth.thrift   
hz update -idl idl/mcpserver.thrift
```

## 生成API文档


```bash
swag init -o docs/swagger
```