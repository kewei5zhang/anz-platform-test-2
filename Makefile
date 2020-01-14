# GO ENV VARIABLES
BUILD_IMAGE=anz-test-2:build
HOST_PORT=8080
CONTAINER_PORT=8000

.PHONY: go-build
go-build: ## Build builder go image
	docker build --target builder -t $(BUILD_IMAGE) .  

.PHONY: go-test
go-test: ## Run go test
	docker run --rm -w /go/src/platform-test/src $(BUILD_IMAGE) go test -v

.PHONY: go-deploy
go-deploy: ## Deploy api
	docker build -t $(BUILD_IMAGE) . 
	docker run --rm -p ${HOST_PORT}:${CONTAINER_PORT} $(BUILD_IMAGE)

##@ Misc
.PHONY: help
help: ## Display help
	@awk \
	  'BEGIN { \
	    FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n" \
	  } /^[a-zA-Z_-]+:.*?##/ { \
	    printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 \
	  } /^##@/ { \
	    printf "\n\033[1m%s\033[0m\n", substr($$0, 5) \
	  }' $(MAKEFILE_LIST)	