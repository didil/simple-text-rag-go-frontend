
protoc-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


gen-protos:
	protoc --go_out=grpc_gen/ --go_opt=paths=source_relative \
    --go-grpc_out=grpc_gen/ --go-grpc_opt=paths=source_relative\
    protos/qa.proto