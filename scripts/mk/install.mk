# Code generated by craft; DO NOT EDIT.

.PHONY: install-golangci-lint
install-golangci-lint:
	@curl -fsSL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | sh -s -- -b "${HOME}/go/bin"

define install_go
current_version=$(go version || echo "go0.0.0")
new_version=$(curl -fsSL "https://go.dev/dl/?mode=json" | jq -r '.[0].version')
if echo "${current_version}" | grep -Eq "${new_version}"; then
	echo "latest go version ${new_version} already installed"
	exit 0
fi

echo "installing latest go version ${new_version}"
rm -rf "${HOME}/.local/go" && mkdir -p "${HOME}/.local/go"
curl -fsSL "https://go.dev/dl/${new_version}.linux-amd64.tar.gz" | (cd "${HOME}/.local/go" && tar -xz --strip-components=1)
for item in "go" "gofmt"; do
	chmod +x "${HOME}/.local/go/bin/${item}" && ln -sf "${HOME}/.local/go/bin/${item}" "${HOME}/.local/bin/${item}"
done
endef
.PHONY: install-go
install-go: ; @$(value install_go)
.ONESHELL:
