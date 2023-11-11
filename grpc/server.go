package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "path/to/proto" // Import generated protobuf code
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
  return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", ":50051")

  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, &GreeterServer{})

  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}