# GitHub Actions 自动部署

生产环境由 `.github/workflows/deploy.yml` 负责。每次代码推送到 `dev` 分支时，流水线先测试并构建两个 Go 服务和两个 Vue 前端；全部通过后，再通过 SSH 连接服务器，执行 `deploy/deploy.sh`。

## 服务器约定

- 仓库目录：`/root/yigeyingshi`
- 部署分支：`dev`
- Compose 文件：`compose.prod.yml`
- 公开站点：端口 `80`
- 公开 API：端口 `8080`，同时可经站点的 `/api/` 访问
- 管理站点：端口 `4174`
- 管理 API：端口 `8081`，同时可经管理站点的 `/api/` 访问
- 公共服务配置：`/root/yigeyingshi/yige-server/.env`
- 管理服务配置：`/root/yigeyingshi/yige-admin-server/.env`

`.env` 文件只保存在服务器上，不进入 Git 仓库，也不会被构建进镜像。

## GitHub production 环境 Secrets

需要在仓库 `Settings → Environments → production` 中配置：

| Secret | 用途 |
| --- | --- |
| `DEPLOY_HOST` | 部署服务器地址 |
| `DEPLOY_PORT` | SSH 端口，通常为 `22` |
| `DEPLOY_USER` | SSH 用户 |
| `DEPLOY_SSH_KEY` | 专用于流水线的 Ed25519 私钥 |
| `DEPLOY_KNOWN_HOSTS` | 服务器 SSH host key，防止连接到错误主机 |

## 手动部署与排查

```sh
cd /root/yigeyingshi
APP_DIR=/root/yigeyingshi DEPLOY_BRANCH=dev sh deploy/deploy.sh <40 位提交 SHA>
docker compose -f compose.prod.yml ps
docker compose -f compose.prod.yml logs --tail=200
```

GitHub Actions 也支持从 Actions 页面使用 `workflow_dispatch` 手动运行。相同环境下的部署会串行执行，避免多个提交同时更新服务器。

每个版本使用提交 SHA 作为 Docker 镜像标签。服务器会在 `/root/.yigeyingshi-current-tag` 记录当前成功版本；如果新容器启动失败或健康检查超时，部署脚本会自动切回上一个成功版本。
