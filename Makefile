CLI_SUFFIX=
DENO_VERSION := 1.22.3

ifeq ($(OS),Windows_NT)
    HOST=windows
    CLI_SUFFIX=.exe
    DENO_INSTALL = build\cli
else
    HOST_S := $(shell uname -s)
    DENO_INSTALL = build/cli
    ifeq ($(HOST_S),Linux)
        HOST=linux

    endif
    ifeq ($(HOST_S),Darwin)
        HOST +=macos
    endif
    HOST_P := $(shell uname -p)
    ifeq ($(HOST_P),x86_64)
        ARCH=amd64
    endif
    ifneq ($(filter %86,$(HOST_P)),)
        ARCH=i386
    endif
    ifneq ($(filter arm%,$(HOST_P)),)
        ARCH=arm
    endif
endif

include deno.mk

.PHONY: all setup help clean build dev
all: help

dev: ## Run dev build
	wails dev

build: ## make prod build
	wails build --clean -webview2 embed

setup: ## Get Build Dependencies
	(cd frontend && npm install --legacy-peer-deps)
	(cd server && npm install)
	npm install

clean: ## removes all the installables
	rm -rf ./node_nodules
	rm -rf ./frontend/node_modules
	rm -rf ./server/node_modules
	rm -rf **/package.json.md5

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-30s\033[0m %s\n", $$1, $$2}'