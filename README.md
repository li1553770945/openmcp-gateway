# openmcp 端口转发工具-后端

## QuickStart

如果你不想参与开发，只是想要部署和使用本项目，请直接参考[QuickStart](./docs/QuickStart.md)

## API 文档

所有的API均会返回code, message两字段，当code不为0时，表示请求失败，message中会包含失败原因。根据API的不同，是否会返回data字段，以及成功时返回的数据结构会有所不同，具体请参考API说明。

code含义如下：

| code         | 含义             |
| ------------ | ---------------- |
| 0            | 成功             |
| 4001         | 参数错误         |
| 4003         | 未授权           |
| 4004         | 资源未找到       |
| 5001         | 系统错误         |

请在部署后访问`http://<server地址>:<server端口>/docs`查看完整的接口说明。

## 开发指南

我们欢迎任何形式的贡献。如果你想参与开发，请参考[开发指南](./docs/DevelopmentGuide.md)了解更多信息。