export

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ## Build and run the server with Docker
	@VER=custom; \
	if [ -f .git/refs/heads/master ]; then VER=$$(head -1 .git/refs/heads/master | head -c 7); fi; \
	sed -i "s/VERSIONSTRING = \".*\"/VERSIONSTRING = \"$$VER\"/g" ./src/config/config.go; \
	cd ./src; \
	go mod tidy && go mod download && \
	docker run --rm -it $$(docker build -q .)
.PHONY: run
