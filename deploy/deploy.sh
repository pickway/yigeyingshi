#!/usr/bin/env sh
set -eu

APP_DIR="${APP_DIR:-/root/yigeyingshi}"
DEPLOY_BRANCH="${DEPLOY_BRANCH:-dev}"
DEPLOY_SHA="${1:-}"
COMPOSE_FILE="${APP_DIR}/compose.prod.yml"
STATE_FILE="${DEPLOY_STATE_FILE:-/root/.yigeyingshi-current-tag}"

case "${DEPLOY_SHA}" in
  ''|*[!0-9a-f]*)
    echo "错误：必须传入完整的 Git 提交 SHA" >&2
    exit 2
    ;;
esac

if [ "${#DEPLOY_SHA}" -ne 40 ]; then
  echo "错误：Git 提交 SHA 长度必须为 40" >&2
  exit 2
fi

cd "${APP_DIR}"

echo "[deploy] 获取 ${DEPLOY_BRANCH} / ${DEPLOY_SHA}"
git fetch --prune origin "${DEPLOY_BRANCH}"
git cat-file -e "${DEPLOY_SHA}^{commit}"
git checkout "${DEPLOY_BRANCH}"
git merge --ff-only "${DEPLOY_SHA}"

if [ ! -f "${APP_DIR}/yige-server/.env" ]; then
  echo "错误：缺少 yige-server/.env" >&2
  exit 3
fi

if [ ! -f "${APP_DIR}/yige-admin-server/.env" ]; then
  echo "错误：缺少 yige-admin-server/.env" >&2
  exit 3
fi

mkdir -p /var/log/yige /var/log/yige-admin
chown -R 100:101 /var/log/yige /var/log/yige-admin

previous_tag=""
if [ -f "${STATE_FILE}" ]; then
  previous_tag="$(sed -n '1p' "${STATE_FILE}")"
fi

case "${previous_tag}" in
  ''|*[!0-9a-f]*) previous_tag="" ;;
esac

echo "[deploy] 构建四个应用镜像"
IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" build --pull

echo "[deploy] 更新服务"
if ! IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" up -d --remove-orphans; then
  echo "错误：新版本容器启动失败" >&2
  IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" ps >&2 || true
  IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" logs --tail=120 >&2 || true
  if [ -n "${previous_tag}" ] && [ "${previous_tag}" != "${DEPLOY_SHA}" ]; then
    echo "[deploy] 回滚到 ${previous_tag}" >&2
    IMAGE_TAG="${previous_tag}" docker compose -f "${COMPOSE_FILE}" up -d --remove-orphans || true
  fi
  exit 4
fi

echo "[deploy] 等待健康检查"
attempt=0
until curl --fail --silent --show-error --output /dev/null http://127.0.0.1/ \
  && curl --fail --silent --show-error --output /dev/null http://127.0.0.1/api/v1/movies \
  && curl --fail --silent --show-error --output /dev/null http://127.0.0.1:4174/ \
  && curl --fail --silent --show-error --output /dev/null http://127.0.0.1:8081/health; do
  attempt=$((attempt + 1))
  if [ "${attempt}" -ge 18 ]; then
    echo "错误：服务健康检查超时" >&2
    IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" ps >&2 || true
    IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" logs --tail=120 >&2 || true
    if [ -n "${previous_tag}" ] && [ "${previous_tag}" != "${DEPLOY_SHA}" ]; then
      echo "[deploy] 回滚到 ${previous_tag}" >&2
      IMAGE_TAG="${previous_tag}" docker compose -f "${COMPOSE_FILE}" up -d --remove-orphans || true
    fi
    exit 5
  fi
  sleep 5
done

state_tmp="${STATE_FILE}.tmp"
printf '%s\n' "${DEPLOY_SHA}" > "${state_tmp}"
chmod 600 "${state_tmp}"
mv -f "${state_tmp}" "${STATE_FILE}"

docker image prune -f --filter "until=168h" >/dev/null
IMAGE_TAG="${DEPLOY_SHA}" docker compose -f "${COMPOSE_FILE}" ps
echo "[deploy] 部署完成：${DEPLOY_SHA}"
