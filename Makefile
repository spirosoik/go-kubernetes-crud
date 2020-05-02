.PHONY: example

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Clean
	go clean

example: ## Run the example. Provide KUBE_CONTEXT env variable
	go run examples/example.go -kubecontext=${KUBE_CONTEXT} 

lint: ## Linting the codebase
	golint -set_exit_status ./...

setup: ## Setup modules
	go get -u golang.org/x/lint/golint