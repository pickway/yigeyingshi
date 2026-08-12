# CineVerse 管理后台 API

独立的 Go 管理服务，使用 Gin、GORM 和 SQLite，为 CineVerse 管理端提供登录、统计和内容 CRUD 接口。默认连接现有的 `../yige-server/yige.db`，因此后台保存的数据会直接进入前台网站的数据源。

## 启动

```powershell
Copy-Item .env.example .env
go run ./cmd/server
```

默认地址为 `http://127.0.0.1:8081`，健康检查为 `/health`。开发默认账号是 `admin / change-me`，请在 `.env` 中立即修改。生产模式还要求 `ADMIN_TOKEN_SECRET` 至少为 32 位。

## 验证

```powershell
go test ./...
go vet ./...
```

所有 `/api/admin` 管理接口（登录除外）都需要 `Authorization: Bearer <token>`。令牌默认 8 小时过期。
