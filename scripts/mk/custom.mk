.PHONY: testdata
testdata: build
	@find . -name go.mod -execdir go mod tidy \;
	@find . -name go.mod -execdir go generate ./... \;
	@echo "It's expected to have 'no such tool' error from generated files, generation for those is still made"

update:
	@./scripts/sh/update.sh
