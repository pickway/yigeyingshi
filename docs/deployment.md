# GitHub Actions 自动部署

生产环境由 `.github/workflows/deploy.yml` 负责。每次代码推送到 `dev` 分支时，流水线先根据改动目录识别受影响服务，只测试、构建和部署这些服务，再通过 SSH 将目标提交直接传到服务器并执行 `deploy/deploy.sh`。服务器无需直接访问 GitHub。

## 目录与服务映射

| 改动路径 | 测试和部署范围 |
| --- | --- |
| `yige-server/**` | 公共 API |
| `yige-ui/**` | 公共站点 |
| `yige-admin-server/**` | 管理 API |
| `yige-admin-ui/**` | 管理站点 |
| `scripts/**` | 两个 Go 后端 |
| `deploy/**`、`compose.prod.yml`、部署工作流 | 全部四个服务 |
| 文档及其他路径 | 不触发测试和部署 |

在 Actions 页面手动运行工作流时，会测试并部署全部四个服务。目录映射由 `deploy/detect-services.sh` 统一维护，并由 `deploy/detect-services.test.sh` 自动验证。

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
git fetch --prune origin dev
APP_DIR=/root/yigeyingshi DEPLOY_BRANCH=dev sh deploy/deploy.sh <40 位提交 SHA> yige-ui
# 也可以一次指定多个服务
APP_DIR=/root/yigeyingshi DEPLOY_BRANCH=dev sh deploy/deploy.sh <40 位提交 SHA> yige-server yige-ui
docker compose -f compose.prod.yml ps
docker compose -f compose.prod.yml logs --tail=200
```

GitHub Actions 也支持从 Actions 页面使用 `workflow_dispatch` 手动运行。相同环境下的部署会串行执行，避免多个提交同时更新服务器。

每个版本使用提交 SHA 作为 Docker 镜像标签。服务器会在 `/root/.yigeyingshi-deploy-state/` 中分别记录四个服务当前的成功版本；如果新容器启动失败或健康检查超时，部署脚本只回滚本次受影响的服务。首次升级时，脚本会自动兼容原有的 `/root/.yigeyingshi-current-tag` 状态文件。
