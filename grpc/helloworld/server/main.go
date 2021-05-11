package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	api "github.com/mrhjkim/gostudy/grpc/helloworld/api/v1"
)

const (
	port = ":50051"
)

type server struct {
	api.UnimplementedHelloworldServer
}

func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %s", in.GetName())
	return &api.HelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// protobuf impl need grpc.Server and api impl
	api.RegisterHelloworldServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
