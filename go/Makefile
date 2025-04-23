.PHONY: install
install-macos-linux:
	@echo "Installing dependencies"
	brew install bufbuild/buf/buf
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@echo "Done"

.PHONY: generate
generate:
	@echo "Generating go bindings"
	cd api/proto && buf generate
	@echo "Done"

.PHONY: build-node
build-node:
	@echo "Building node"
	go build -o bin/node node/cmd/main.go
	@echo "Done"

.PHONY: fmt
fmt: ## formats all go files
	go fmt ./...

.PHONY: test
tests:
	go test ./...
