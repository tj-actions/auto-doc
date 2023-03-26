# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

OUTPUT := README.md
FILENAME := action.yml

.DEFAULT_GOAL := help

.PHONY: help
# Put it first so that "make" without argument is like "make help".
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s-\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean:  ## Clean binary file
	@echo "Cleaning binary..."
	@rm -rf bin

guard-%: ## Checks that env var is set else exits with non 0 mainly used in CI;
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: build
build: clean  ## Compile go modules
	@echo "Compiling *.go..."
	@go build -o ./bin/auto_doc *.go

.PHONY: run
run: build guard-OUTPUT guard-FILENAME  ## Execute binary
	@echo "Running auto doc..."
	@./bin/auto_doc --filename=$(FILENAME) --output=$(OUTPUT)
	@$(MAKE) clean

.PHONY: run-help
run-help: build guard-OUTPUT guard-FILENAME  ## Execute binary
	@echo "Running auto doc help..."
	@echo ""
	@./bin/auto_doc --help
	@$(MAKE) clean

upgrade-major-version:  ## Upgrade major version
	@find . -type f \
        -name '*.go' \
        -exec sed -i 's,github.com/tj-actions/auto-doc,github.com/tj-actions/auto-doc/v2,g' {} \;

.PHONY: test
test: clean
	@go test ./cmd

.PHONY: format
format:  ## Format go modules
	@go fmt ./...

.PHONY: tidy
tidy:  ## Tidy go.mod
	@go mod tidy
