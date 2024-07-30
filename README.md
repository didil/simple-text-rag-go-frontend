
# Development Setup

## Install protoc
```bash
brew install protobuf
```

## Install the protocol compiler plugins for Go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Setup .env file
```bash
cp .env.example .env
```