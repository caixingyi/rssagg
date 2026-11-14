# rssagg

一个用 Go 编写的轻量 RSS 聚合器。用户可以创建订阅源（feed）、关注订阅源，后台抓取器会定期抓取这些订阅源并存储文章，用户再通过 JSON API 阅读。

## 技术栈

- Go 1.26
- PostgreSQL
- [go-chi](https://github.com/go-chi/chi) 路由
- [goose](https://github.com/pressly/goose) 数据库迁移
- [sqlc](https://sqlc.dev) 类型安全的 SQL 代码生成

## 前置要求

- Go 1.26+
- 本地已运行 PostgreSQL
- 已安装 `goose` 和 `sqlc`：
  ```sh
  go install github.com/pressly/goose/v3/cmd/goose@latest
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  ```

## 安装与启动

1. 创建数据库：
   ```sh
   createdb rssagg
   ```

2. 配置环境变量（见 `.env`）：
   ```sh
   PORT=8000
   DB_URL=postgres://user:password@localhost:5432/rssagg?sslmode=disable
   ```

3. 执行数据库迁移：
   ```sh
   goose -dir sql/schema postgres "$DB_URL" up
   ```

4. 修改查询或表结构后，重新生成 sqlc 代码：
   ```sh
   sqlc generate
   ```

5. 启动服务：
   ```sh
   go run .
   ```

## API

所有需要鉴权的接口，都要带上 `Authorization: ApiKey <key>` 请求头。

| 方法 | 路径 | 鉴权 | 说明 |
|------|------|------|------|
| POST | `/v1/users` | 否 | 创建用户，返回 `api_key` |
| GET | `/v1/users` | 是 | 获取当前用户 |
| POST | `/v1/feeds` | 是 | 创建订阅源 |
| GET | `/v1/feeds` | 否 | 列出所有订阅源 |
| POST | `/v1/feed_follows` | 是 | 关注某个订阅源 |
| GET | `/v1/feed_follows` | 是 | 列出我关注的订阅源 |
| DELETE | `/v1/feed_follows/{id}` | 是 | 取消关注 |
| GET | `/v1/posts` | 是 | 列出我关注的订阅源里的文章 |
| GET | `/v1/health` | 否 | 健康检查 |
| GET | `/v1/err` | 否 | 调试用接口 |

## 抓取机制

`startScraping` 在后台运行，通过 ticker 定时、并发地抓取订阅源。收到 SIGINT/SIGTERM（例如 Ctrl+C）时会优雅停机。

## 测试

```sh
go test ./...
```
