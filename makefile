
PROTO_DIR := api

PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)


GO_OUT := pkg/pb

.PHONY: all proto test lint build clean run

all: build


proto:
	protoc -I api/ --go_out=. --go_opt=paths=source_relative \
  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
  	--validate_out="lang=go:." api/file.proto



lint:
	golangci-lint run

test:
	./test/grpc_load_test.sh localhost:50051 pb.FileService/UploadFile 10000 50 payload.json

build:
	mkdir -p bin
	go build -o bin/server ./cmd/server


clean:
	rm -rf bin

run: build
	./bin/server

