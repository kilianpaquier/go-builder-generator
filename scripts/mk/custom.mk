.PHONY: go-generate
go-generate:
	@go generate ./...
	@go generate ./testdata/**
	@find ./testdata -name go.mod -execdir go mod tidy \;
	@find ./testdata -name go.mod -execdir go generate ./... \;
