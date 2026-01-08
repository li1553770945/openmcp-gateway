# 配置说明


```yaml
server: # 服务器配置
  listen-address: 0.0.0.0:9000 # 服务器监听地址
  proxy-listen-address: 0.0.0.0:9001 # 代理服务器监听地址

database: # 数据库配置
  username: your_db_username # 数据库用户名
  password: your_db_password # 数据库密码
  database: your_database_name # 数据库名称
  address: your_database_address # 数据库地址
  port: 21024 # 数据库端口
  use-tls: false # 是否使用 TLS

auth:
  jwt-key: "your_jwt_key" # 用于 JWT 认证的密钥，请使用强随机字符串且务必保密，更新此密钥会使所有现有的 JWT 失效

```