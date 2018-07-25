
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=. --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --micro_out=. --go_out=plugins=grpc:. proto/micro/auth/auth.proto
	protoc --proto_path=. --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. proto/rest/auth/auth.proto
	protoc --proto_path=. --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. proto/rest/auth/auth.proto

build: proto

	go build -o auth-api main.go plugin.go

test:
	go test -v ./... -cover

docker:
    GOOS=linux GOARCH=amd64 go build -o auth-api main.go plugin.go
	docker build . -t zergwangj/auth-api:latest
