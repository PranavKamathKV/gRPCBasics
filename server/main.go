package main

import(
	"gRPCBasics/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"context"

)

// To implement the service interface in pb.go
type server struct{}

func main(){
	listener, err:= net.Listen("tcp", ":4040")
	if err != nil{
		panic(err)
	}

	src:= grpc.NewServer()
	proto.RegisterAddServiceServer(src, &server{})
	reflection.Register(src)

	if err:= src.Serve(listener); err!=nil{
		panic(err)
	}

}


func  (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a, b:= request.GetA(), request.GetB()
	c:= a + b

	return &proto.Response{C:c}, nil
}


func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a,b:= request.GetA(), request.GetB()

	c:= a*b

	return &proto.Response{C:c}, nil
}


