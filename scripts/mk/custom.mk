.PHONY: testdata
testdata: build ## regenerate testdata files.
	@find . -name go.mod -execdir go mod tidy \;
	@find . -name go.mod -execdir go generate ./... \;
	@echo "It's expected to have 'no such tool' error from generated files, generation for those is still made"

update: ## update all go.mod dependencies.
	@./scripts/sh/update.sh
