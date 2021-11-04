# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

OUTPUT := README.md
ACTION := action.yml

.DEFAULT_GOAL := help

.PHONY: help
# Put it first so that "make" without argument is like "make help".
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s-\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean:  ## Clean binary file
	@echo "Cleaning binary..."
	@rm -f auto_doc

guard-%: ## Checks that env var is set else exits with non 0 mainly used in CI;
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: build
build:  ## Compile go modules
	@echo "Compiling *.go..."
	@go build -o auto_doc *.go

.PHONY: run
run: build guard-OUTPUT guard-ACTION  ## Execute binary
	@echo "Running auto doc..."
	@./auto_doc --action=$(ACTION) --output=$(OUTPUT)
	@$(MAKE) clean

.PHONY: format
format:  ## Format go modules
	@go fmt ./...

.PHONY: tidy
tidy:  ## Tidy go.mod
	@go mod tidy
