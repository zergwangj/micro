
proto:
	protoc --proto_path=. --proto_path=%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --micro_out=. --go_out=plugins=grpc:. proto/users/users.proto

build: proto

	go build -o users-srv main.go plugin.go

test:
	go test -v ./... -cover

docker:
    GOOS=linux GOARCH=amd64 go build -o users-srv main.go plugin.go
	docker build . -t zergwangj/users-srv:latest
