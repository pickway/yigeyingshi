#!/usr/bin/env sh
set -eu

APP_DIR="${APP_DIR:-/root/yigeyingshi}"
DEPLOY_BRANCH="${DEPLOY_BRANCH:-dev}"
DEPLOY_SHA="${1:-}"
COMPOSE_FILE="${APP_DIR}/compose.prod.yml"
STATE_DIR="${DEPLOY_STATE_DIR:-/root/.yigeyingshi-deploy-state}"
LEGACY_STATE_FILE="${DEPLOY_STATE_FILE:-/root/.yigeyingshi-current-tag}"

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

shift
if [ "$#" -eq 0 ]; then
  echo "错误：至少指定一个待部署服务" >&2
  exit 2
fi

deploy_yige_server=false
deploy_yige_ui=false
deploy_yige_admin_server=false
deploy_yige_admin_ui=false

for service in "$@"; do
  case "${service}" in
    yige-server) deploy_yige_server=true ;;
    yige-ui) deploy_yige_ui=true ;;
    yige-admin-server) deploy_yige_admin_server=true ;;
    yige-admin-ui) deploy_yige_admin_ui=true ;;
    *)
      echo "错误：未知服务 ${service}" >&2
      exit 2
      ;;
  esac
done

set --
if [ "${deploy_yige_server}" = "true" ]; then set -- "$@" yige-server; fi
if [ "${deploy_yige_ui}" = "true" ]; then set -- "$@" yige-ui; fi
if [ "${deploy_yige_admin_server}" = "true" ]; then set -- "$@" yige-admin-server; fi
if [ "${deploy_yige_admin_ui}" = "true" ]; then set -- "$@" yige-admin-ui; fi

valid_tag() {
  candidate="$1"
  case "${candidate}" in
    *[!0-9a-f]*|'') return 1 ;;
  esac
  [ "${#candidate}" -eq 40 ]
}

read_previous_tag() {
  service="$1"
  service_state="${STATE_DIR}/${service}"
  candidate=""

  if [ -f "${service_state}" ]; then
    candidate="$(sed -n '1p' "${service_state}")"
  elif [ -f "${LEGACY_STATE_FILE}" ]; then
    candidate="$(sed -n '1p' "${LEGACY_STATE_FILE}")"
  fi

  if valid_tag "${candidate}"; then
    printf '%s\n' "${candidate}"
  else
    printf '%s\n' latest
  fi
}

cd "${APP_DIR}"

echo "[deploy] 准备 ${DEPLOY_BRANCH} / ${DEPLOY_SHA}"
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

mkdir -p /var/log/yige /var/log/yige-admin "${STATE_DIR}"
chown -R 100:101 /var/log/yige /var/log/yige-admin
chmod 700 "${STATE_DIR}"

previous_yige_server_tag="$(read_previous_tag yige-server)"
previous_yige_ui_tag="$(read_previous_tag yige-ui)"
previous_yige_admin_server_tag="$(read_previous_tag yige-admin-server)"
previous_yige_admin_ui_tag="$(read_previous_tag yige-admin-ui)"

YIGE_SERVER_TAG="${previous_yige_server_tag}"
YIGE_UI_TAG="${previous_yige_ui_tag}"
YIGE_ADMIN_SERVER_TAG="${previous_yige_admin_server_tag}"
YIGE_ADMIN_UI_TAG="${previous_yige_admin_ui_tag}"

if [ "${deploy_yige_server}" = "true" ]; then YIGE_SERVER_TAG="${DEPLOY_SHA}"; fi
if [ "${deploy_yige_ui}" = "true" ]; then YIGE_UI_TAG="${DEPLOY_SHA}"; fi
if [ "${deploy_yige_admin_server}" = "true" ]; then YIGE_ADMIN_SERVER_TAG="${DEPLOY_SHA}"; fi
if [ "${deploy_yige_admin_ui}" = "true" ]; then YIGE_ADMIN_UI_TAG="${DEPLOY_SHA}"; fi
export YIGE_SERVER_TAG YIGE_UI_TAG YIGE_ADMIN_SERVER_TAG YIGE_ADMIN_UI_TAG

show_diagnostics() {
  docker compose -f "${COMPOSE_FILE}" ps "$@" >&2 || true
  docker compose -f "${COMPOSE_FILE}" logs --tail=120 "$@" >&2 || true
}

restore_previous_tags() {
  if [ "${deploy_yige_server}" = "true" ]; then YIGE_SERVER_TAG="${previous_yige_server_tag}"; fi
  if [ "${deploy_yige_ui}" = "true" ]; then YIGE_UI_TAG="${previous_yige_ui_tag}"; fi
  if [ "${deploy_yige_admin_server}" = "true" ]; then YIGE_ADMIN_SERVER_TAG="${previous_yige_admin_server_tag}"; fi
  if [ "${deploy_yige_admin_ui}" = "true" ]; then YIGE_ADMIN_UI_TAG="${previous_yige_admin_ui_tag}"; fi
  export YIGE_SERVER_TAG YIGE_UI_TAG YIGE_ADMIN_SERVER_TAG YIGE_ADMIN_UI_TAG

  echo "[deploy] 回滚受影响服务" >&2
  docker compose -f "${COMPOSE_FILE}" up -d "$@" || true
}

check_selected_services() {
  for service in "$@"; do
    case "${service}" in
      yige-server) url="http://127.0.0.1:8080/health" ;;
      yige-ui) url="http://127.0.0.1/" ;;
      yige-admin-server) url="http://127.0.0.1:8081/health" ;;
      yige-admin-ui) url="http://127.0.0.1:4174/" ;;
    esac
    curl --fail --silent --show-error --output /dev/null "${url}" || return 1
  done
}

echo "[deploy] 构建服务：$*"
docker compose -f "${COMPOSE_FILE}" build --pull "$@"

echo "[deploy] 更新服务：$*"
if ! docker compose -f "${COMPOSE_FILE}" up -d "$@"; then
  echo "错误：新版本容器启动失败" >&2
  show_diagnostics "$@"
  restore_previous_tags "$@"
  exit 4
fi

echo "[deploy] 等待健康检查"
attempt=0
until check_selected_services "$@"; do
  attempt=$((attempt + 1))
  if [ "${attempt}" -ge 18 ]; then
    echo "错误：服务健康检查超时" >&2
    show_diagnostics "$@"
    restore_previous_tags "$@"
    exit 5
  fi
  sleep 5
done

for service in "$@"; do
  state_file="${STATE_DIR}/${service}"
  state_tmp="${state_file}.tmp"
  printf '%s\n' "${DEPLOY_SHA}" > "${state_tmp}"
  chmod 600 "${state_tmp}"
  mv -f "${state_tmp}" "${state_file}"
done

docker image prune -f --filter "until=168h" >/dev/null
docker compose -f "${COMPOSE_FILE}" ps "$@"
echo "[deploy] 部署完成：${DEPLOY_SHA}（$*）"
