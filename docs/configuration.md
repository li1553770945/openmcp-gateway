# 配置说明 (Configuration)

本项目使用 YAML 格式进行配置。根据运行环境不同，系统会加载 `conf/deployment.yml` (开发) 或 `conf/production.yml` (生产)。

## 📝 字段详细说明

### 1. Server (服务器设置)

| 字段 | 类型 | 必填 | 默认值 | 说明 |
| --- | --- | --- | --- | --- |
| `server.listen-address` | String | 是 | `0.0.0.0:9000` | 后端服务的监听地址与端口。 |
### 2. Database (数据库设置)

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `database.type` | String | 是 | 使用的数据库类型，目前支持 `mysql` 和 `sqlite`。 |
| `database.database` | String | 是 | 使用的数据库名称，如果是 SQLite，则为文件路径。 |
| `database.username` | String | 是 | 数据库登录用户名。 |
| `database.password` | String | 是 | 数据库登录密码。 |
| `database.address` | String | 是 | 数据库连接地址（不带端口）。 |
| `database.port` | Int | 是 | 数据库端口（如 MySQL 默认 3306）。 |
| `database.use-tls` | Boolean | 否 | 是否启用 SSL/TLS 加密连接。 |

### 3. Auth (安全认证)

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| `auth.jwt-key` | String | 是 | **核心密钥**。用于签发 Token，建议使用 32 位以上随机字符串。 |

### 4. Proxy (代理设置)
| 字段 | 类型 | 必填 | 默认值 | 说明 |
| --- | --- | --- | --- | --- |
| `proxy.cache-expiration-seconds` | Int | 否 | `60` | 代理缓存过期时间，单位为秒。如果您设置了缓存，将会提高转发性能，但可能导致配置变更后生效有几秒的略微延迟。 |

---

## 🔐 安全建议

> [!CAUTION]
> **切勿泄露您的 `jwt-key`！**
> 在生产环境下，请务必修改默认密钥。如果密钥丢失或被更改，所有已登录的用户令牌（Token）将立即失效。

---

## 📄 完整配置示例

请参考项目根目录下的`config-example.yml`文件获取完整的配置示例。
