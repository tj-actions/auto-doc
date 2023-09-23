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

upgrade-from-one-version-to-a-major-version: guard-OLD_VERSION guard-MAJOR_VERSION  ## Upgrade from and old version to a major version
	@echo "Upgrading from v$(OLD_VERSION) to v$(MAJOR_VERSION)"
	@find . -type f \
		-name '*.go' \
		-exec sed -i '' -e 's,github.com/tj-actions/auto-doc/v$(OLD_VERSION),github.com/tj-actions/auto-doc/v$(MAJOR_VERSION),g' {} \;
	@find . -type f \
		-name 'go.mod' \
		-exec sed -i '' -e 's,github.com/tj-actions/auto-doc/v$(OLD_VERSION),github.com/tj-actions/auto-doc/v$(MAJOR_VERSION),g' {} \;
	@find . -type f \
		-name 'README.md' \
		-exec sed -i '' -e 's,github.com/tj-actions/auto-doc/v$(OLD_VERSION),github.com/tj-actions/auto-doc/v$(MAJOR_VERSION),g' {} \;

.PHONY: test
test: clean
	@go test ./cmd

.PHONY: format
format:  ## Format go modules
	@go fmt ./...

.PHONY: tidy
tidy:  ## Tidy go.mod
	@go mod tidy
