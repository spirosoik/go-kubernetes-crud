help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## Linting the codebase
	golint -set_exit_status ./...

setup: ## Setup modules
	go get -u golang.org/x/lint/golint

clean: ## Clean
	go clean