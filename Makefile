.PHONY: proto
proto: ## Generate the Go code from our protobuf definitions
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative analyser.proto

test: ## Run tests
	go test -v ./...

server-build: ## Build the server
	cd service && go build -o ../bin/server

server-run: ## Run the server
	cd service && go run .

cli-build: ## Build the cli
	go build -o bin/analyser

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'