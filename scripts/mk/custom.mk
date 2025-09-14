.PHONY: testdata
testdata:
	@go build -o go-builder-generator ./cmd/go-builder-generator
	@find . -name go.mod -execdir go mod tidy \;
	@find . -name go.mod -execdir go generate ./... \;
	@echo "It's expected to have the error 'no such tool' from generated files"
