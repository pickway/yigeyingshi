# CineVerse / 一个影视

一套围绕电影、影视创作学习与 AI 工具观察构建的中文个人内容网站。项目包含 Vue 单页应用、Go JSON API 和 SQLite 内容库。

## 功能

- 首页个人品牌叙事、本周放映、近期文章与 AI 工具精选
- 电影关键词搜索、类型筛选、排序、分页与详情笔记
- 学习课程分类、推荐路径与本地学习进度交互
- AI 工具目录、相关文章与邮箱订阅
- 全站搜索：同时检索影片、文章、课程和工具
- 关于页、404 页面、移动端导航，以及完整加载/空/错状态
- 独立管理后台：安全登录、内容总览、五类内容 CRUD 与订阅用户管理

## 本地运行

需要 Node.js 20+、npm 和 Go 1.22+。

PowerShell 中，按仓库规则先清空代理变量：

```powershell
$proxyNames = @('HTTP_PROXY','HTTPS_PROXY','ALL_PROXY','NO_PROXY','http_proxy','https_proxy','all_proxy','no_proxy')
foreach ($proxyName in $proxyNames) { Remove-Item -Path "Env:$proxyName" -ErrorAction SilentlyContinue }
```

后端：

```powershell
cd yige-server
Copy-Item .env.example .env
go run ./cmd/server
```

前端（另一个终端）：

```powershell
cd yige-ui
npm install
npm run dev
```

打开 Vite 输出的本地地址。开发服务器会把 `/api` 代理到 `http://localhost:8080`。

管理后台使用两个独立目录。先启动管理 API：

```powershell
cd yige-admin-server
Copy-Item .env.example .env
go run ./cmd/server
```

再启动管理前端：

```powershell
cd yige-admin-ui
npm install
npm run dev
```

打开 `http://127.0.0.1:4174`。开发默认账号为 `admin / change-me`，正式使用前必须在 `yige-admin-server/.env` 中更换密码和令牌密钥。

## 验证

```powershell
cd yige-ui
npm test
npm run build

cd ..\yige-server
go test ./...
go vet ./...
```

## 配置与数据

后端默认监听 `0.0.0.0:8080`，数据库文件为 `yige-server/yige.db`。可通过 `.env` 修改：

```dotenv
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
GIN_MODE=debug
DB_DRIVER=sqlite
DB_DSN=./yige.db
```

首次运行会自动迁移表结构并写入演示内容。数据库文件、`.env`、前端构建目录和依赖目录均已忽略，不进入版本控制。

## 目录

```text
yige-ui/                 Vue 3 + Vite 前端
yige-server/             Gin + GORM + SQLite 后端
yige-admin-ui/           Vue 3 + Element Plus 管理前端
yige-admin-server/       独立的 Go 管理 API
cinema-personal-site/    早期视觉原型，保留作设计参考
docs/                    产品规格与实施方案
```

生产环境应由反向代理提供 HTTPS、静态文件缓存与 `/api` 转发，并通过可信域名限制 CORS。当前匿名写接口只有订阅，若公开部署建议在网关增加限流或验证码。
