
.PHONY: all setup help clean build dev server
all: help

dev: ## Run dev build
	wails dev

server-linux:
	(cd server && deno task build-amd-linux && deno task build-arm-linux)

server-windows:
	(cd server && deno task build-amd-win)

server-macos:
	(cd server && deno task build-amd-macos && deno task build-arm-macos)


build: ## make prod build
	wails build --clean -webview2 embed

setup: ## Get Build Dependencies
	(cd frontend && npm install --legacy-peer-deps)
	npm install

clean: ## removes all the installables
	rm -rf ./node_nodules
	rm -rf ./frontend/node_modules
	rm -rf **/package.json.md5

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-30s\033[0m %s\n", $$1, $$2}'
