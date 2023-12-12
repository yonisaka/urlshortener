GOPATH=$(shell go env GOPATH)

GOLANGCI_LINT_VERSION=v1.51.1

install: install-gomock install-linter

install-gomock:
	@echo "\n>>> Install gomock\n"
	go install github.com/golang/mock/mockgen

install-linter:
	@echo "\n>>> Install GolangCI-Lint"
	@echo ">>> https://github.com/golangci/golangci-lint/releases \n"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_LINT_VERSION}/install.sh | \
	sh -s -- -b ${GOPATH}/bin ${GOLANGCI_LINT_VERSION}

	@echo "\n>>> Install protolint"
	@echo ">>> https://github.com/yoheimuta/protolint/releases \n"
	@go install github.com/yoheimuta/protolint/cmd/protolint

lint:
	@echo "\n>>> Run GolangCI-Lint\n"
	/bin/bash ./scripts/lint.sh

	@echo "\n>>> Run Proto-Lint\n"
	protolint api/proto

test:
	mkdir -p .coverage/html
	go test -v -race -cover -coverprofile=.coverage/internal.coverage ./internal/... && \
	cat .coverage/internal.coverage | grep -v "_mock.go\|_mockgen.go" > .coverage/internal.mockless.coverage && \
	go tool cover -html=.coverage/internal.mockless.coverage -o .coverage/html/internal.coverage.html;

diagram:
	@echo "\n>>> Run Generate Diagram\n"
	go run ./scripts/generate-diagram

sequence-diagram:
	@echo "\n>>> Run Generate RPC Sequence Diagram\n"
	go run ./scripts/generate-rpc-sequence-diagram -RPC=$(RPC)

protoc:
	@echo "\n>>> Run Generate Protoc\n"
	buf generate
	buf generate --template buf.gen-apis.yaml --path api/proto/service.proto

mock:
	@echo "\n>>> Run Generate Mock\n"
	go generate ./...

build-protoc:
	docker build -t urlshortener-protoc -f build/protoc/Dockerfile .

docker-protoc:
	docker run --rm -v `pwd`:/go/src/github.com/yonisaka/urlshortener urlshortener-protoc
