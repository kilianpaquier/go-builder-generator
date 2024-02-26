# This file is safe to edit. Once it exists it will not be overwritten.

include ./scripts/*.mk

.PHONY: generate
generate:
	@craft generate $(ARGS)

.PHONY: clean
clean:
	@go clean
	@git clean -Xf ./*