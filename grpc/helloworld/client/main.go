package main

import (
	"context"
	"log"

	api "github.com/mrhjkim/gostudy/grpc/helloworld/api/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// protobuf Client impl need grpc.ClientConn
	c := api.NewHelloworldClient(conn)
	ctx := context.Background()
	r, err := c.SayHello(ctx, &api.HelloRequest{
		Name: "mrhjkim",
	})
	if err != nil {
		log.Fatalf("fail SayHello : %v", err)
	}
	log.Printf("Greeting : %s", r.GetMessage())
}
