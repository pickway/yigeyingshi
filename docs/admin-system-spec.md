# Spec: CineVerse 后台管理系统

## Objective

新建一套与公开站解耦的后台管理系统，供单管理员维护影片、文章、课程、学习路径、AI 工具与订阅用户。后台连接公开站现有 SQLite 数据库，修改后前台接口立即读取新内容。

## Tech Stack

- `yige-admin-server`: Go 1.22, Gin, GORM, SQLite
- `yige-admin-ui`: Vue 3, Vite, Vue Router, Element Plus
- Auth: 环境变量账号密码 + HMAC 签名、短时效 Bearer Token

## Commands

执行前清空大小写代理变量。

- Server dev: `go run ./cmd/server`
- Server test: `go test ./...`
- UI dev: `npm run dev`
- UI test: `npm test`
- UI build: `npm run build`

## Structure

- `yige-admin-server/cmd/server`: 启动入口
- `yige-admin-server/internal/config`: 环境配置
- `yige-admin-server/internal/model`: 与公开数据库兼容的模型
- `yige-admin-server/internal/auth`: 令牌签发与认证中间件
- `yige-admin-server/internal/handler`: 仪表盘和资源 CRUD
- `yige-admin-ui/src/views`: 登录、仪表盘、资源管理页
- `yige-admin-ui/src/components`: 布局与通用编辑器
- `yige-admin-ui/src/api`: 管理 API 客户端

## API Contract

- `POST /api/admin/auth/login`
- `GET /api/admin/dashboard`
- `GET|POST /api/admin/{movies|articles|courses|paths|tools}`
- `GET|PUT|DELETE /api/admin/{resource}/:id`
- `GET /api/admin/subscribers`
- `PATCH /api/admin/subscribers/:id/status`
- `DELETE /api/admin/subscribers/:id`

成功响应使用 `{ data }`，列表额外返回 `{ total, page, pageSize }`。错误统一使用 `{ error: { code, message } }`。

## Code Style

```go
func (h *ResourceHandler) ListMovies(c *gin.Context) {
	h.list(c, &[]model.Movie{}, "movies", []string{"title", "director"})
}
```

方法调用显式使用接收者；输入仅在 HTTP 边界验证；查询使用 GORM 参数绑定。

## Testing Strategy

- 令牌签发、过期和错误签名使用单元测试。
- 登录、未认证拒绝、仪表盘及电影 CRUD 使用内存 SQLite 集成测试。
- 前端纯函数测试 token 与查询序列化；生产构建验证全部 SFC。
- 浏览器验证登录、仪表盘、创建/编辑/删除一条临时内容的完整流程。

## Boundaries

- Always: 所有管理接口鉴权，密码不写入代码或日志，分页有上限，删除二次确认。
- Ask first: 多管理员、角色权限、上传文件、外部对象存储。
- Never: 默认凭据用于生产，不加鉴权暴露写接口，提交 `.env` 或数据库文件。

## Success Criteria

- 新建两个独立目录且均可单独启动、测试和构建。
- 管理员可登录并查看各类内容数量与最近内容。
- 六类业务数据均能在后台查看；五类内容支持 CRUD，订阅者支持状态切换和删除。
- 文章编辑支持正文，课程与路径字段可完整维护。
- 桌面与平板布局可用，错误、空态、加载态完整。
- 后台修改共享数据库后，公开站 API 可读取结果。

