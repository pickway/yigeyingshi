#!/usr/bin/env sh
set -eu

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
DETECTOR="${SCRIPT_DIR}/detect-services.sh"

assert_detection() {
  name="$1"
  paths="$2"
  expected="$3"

  actual="$(printf '%s\n' "${paths}" | sh "${DETECTOR}")"
  if [ "${actual}" != "${expected}" ]; then
    echo "失败：${name}" >&2
    echo "期望：" >&2
    printf '%s\n' "${expected}" >&2
    echo "实际：" >&2
    printf '%s\n' "${actual}" >&2
    exit 1
  fi

  echo "通过：${name}"
}

none='yige_server=false
yige_ui=false
yige_admin_server=false
yige_admin_ui=false
any=false'

public_ui='yige_server=false
yige_ui=true
yige_admin_server=false
yige_admin_ui=false
any=true'

both_uis='yige_server=false
yige_ui=true
yige_admin_server=false
yige_admin_ui=true
any=true'

backends='yige_server=true
yige_ui=false
yige_admin_server=true
yige_admin_ui=false
any=true'

all='yige_server=true
yige_ui=true
yige_admin_server=true
yige_admin_ui=true
any=true'

assert_detection "仅前台页面" "yige-ui/src/App.vue" "${public_ui}"
assert_detection "同时修改两个前端" "yige-ui/package.json
yige-admin-ui/src/views/Login.vue" "${both_uis}"
assert_detection "文档不触发部署" "README.md
docs/deployment.md" "${none}"
assert_detection "数据库脚本影响两个后端" "scripts/cineverse-init.sql" "${backends}"
assert_detection "部署配置影响全部服务" "compose.prod.yml" "${all}"
assert_detection "工作流改动影响全部服务" ".github/workflows/deploy.yml" "${all}"

actual="$(printf '%s\n' 'README.md' | DEPLOY_ALL=true sh "${DETECTOR}")"
if [ "${actual}" != "${all}" ]; then
  echo "失败：手动部署全部服务" >&2
  exit 1
fi
echo "通过：手动部署全部服务"
