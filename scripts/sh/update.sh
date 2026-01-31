#!/bin/sh

log_info() {
  fg="\033[0;34m"
  reset="\033[0m"
  echo "${fg}$1${reset}"
}

dirs=$(find "$(pwd)" -name go.mod -exec dirname {} +;)
for dir in $dirs; do
    log_info "Updating go dependencies in $dir"
    (cd "$dir" && go get -u ./... && go mod tidy)
done
unset dirs dir
