#!/usr/bin/env sh
set -eu

yige_server=false
yige_ui=false
yige_admin_server=false
yige_admin_ui=false

mark_all() {
  yige_server=true
  yige_ui=true
  yige_admin_server=true
  yige_admin_ui=true
}

if [ "${DEPLOY_ALL:-false}" = "true" ]; then
  mark_all
else
  while IFS= read -r changed_path; do
    case "${changed_path}" in
      yige-server/*)
        yige_server=true
        ;;
      yige-ui/*)
        yige_ui=true
        ;;
      yige-admin-server/*)
        yige_admin_server=true
        ;;
      yige-admin-ui/*)
        yige_admin_ui=true
        ;;
      scripts/*)
        yige_server=true
        yige_admin_server=true
        ;;
      .github/workflows/deploy.yml|compose.prod.yml|deploy/*)
        mark_all
        ;;
    esac
  done
fi

any=false
if [ "${yige_server}" = "true" ] \
  || [ "${yige_ui}" = "true" ] \
  || [ "${yige_admin_server}" = "true" ] \
  || [ "${yige_admin_ui}" = "true" ]; then
  any=true
fi

printf '%s\n' \
  "yige_server=${yige_server}" \
  "yige_ui=${yige_ui}" \
  "yige_admin_server=${yige_admin_server}" \
  "yige_admin_ui=${yige_admin_ui}" \
  "any=${any}"
