FROM golang:1.21.5-alpine AS go

RUN apk update && apk add protobuf-dev protobuf && apk add curl

WORKDIR /usr/local/bin

COPY go.mod  go.sum ./
COPY tools .

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install github.com/envoyproxy/protoc-gen-validate
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

ENV BIN="/usr/local/bin"
ENV VERSION="1.14.0"
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o "${BIN}/buf" && chmod +x "${BIN}/buf"

WORKDIR /go/src/github.com/yonisaka/urlshortener

CMD ["sh","scripts/run-protoc.sh"]
