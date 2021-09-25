# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.DEFAULT_GOAL := help

.PHONY: help
# Put it first so that "make" without argument is like "make help".
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s-\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean binary file
	@echo "Cleaning binary..."
	@rm -f auto_doc

.PHONY: build
build: clean ## Compile go modules
	@echo "Compiling *.go..."
	@go build -o auto_doc *.go

.PHONY: run
run: build ## Execute binary
	@echo "Executing binary..."
	@./auto_doc
	@$(MAKE) clean
