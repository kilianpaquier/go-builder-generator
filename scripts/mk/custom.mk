.PHONY: go-generate
go-generate:
	@find . -name go.mod -execdir go get -u ./... \;
	@find . -name go.mod -execdir go mod tidy \;
	@find . -name go.mod -execdir go generate ./... \;
