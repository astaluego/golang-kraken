.PHONY: all generate help contributors
.DEFAULT: default

all: help

generate: ## generate golang files
	@echo "📌 $@"
	@go generate

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

contributors: ## list contributors
	@echo "📌 $@"
	@git log --format='%aN <%aE>' | sort -fu > $@
