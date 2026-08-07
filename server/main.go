package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	// Import the generated pb package
	pb "github.com/PenguinPoweredApps/unary/proto"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the unary RPC method defined in our proto file.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %v", in.GetName())

	// Construct and return the single response
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 1. Listen on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Create a new gRPC server instance
	s := grpc.NewServer()

	// 3. Register our service implementation with the gRPC server
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	// 4. Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
