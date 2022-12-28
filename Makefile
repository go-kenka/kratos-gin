GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find ginproto -name *.proto)
API_PROTO_FILES=$(shell find example/api -name *.proto)
API_PB_FILES=$(shell find example/api -name *pb.go)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go install github.com/go-kenka/kratos-gin/cmd/protoc-gen-go-gin@latest
	go install github.com/favadi/protoc-go-inject-tag@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./example/api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./example/api \
 	       --go-gin_out=paths=source_relative:./example/api \
 	       --openapi_out==paths=source_relative:. --openapi_opt=enum_type=string\
 	       --validate_out=paths=source_relative,lang=go:./example/api \
 	       --go-errors_out=paths=source_relative:./example/api \
		   $(API_PROTO_FILES) \

.PHONY: inject-tag
# generate inject-tag
inject-tag:
	protoc-go-inject-tag -input=$(API_PB_FILES) 
	       
