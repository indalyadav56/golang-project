package grpc

import (
	"context"
	"log"

	pb "path/to/proto"

	"google.golang.org/grpc"
)

func client() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

  client := pb.NewGreeterClient(conn)

  response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "John"})
  
  log.Println(response.Message)
}