# gRPCBasics


Commands to compile proto file
$ protoc --proto_path=`pwd` --go_out=plugins=grpc:`pwd` service.proto


Commands to run the application:
## Run the server program first
$ go build server/main.go
$ go run server/main.go

## Open another terminal to start client
$ go build client/main.go
$ go run client/main.go
