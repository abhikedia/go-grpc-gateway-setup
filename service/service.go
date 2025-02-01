package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/abhikedia/go-grpc-gateway-setup/protogen/golang/yourservice"
)

// server is used to implement the YourServiceServer interface
type server struct {
	golang.UnimplementedYourServiceServer
}

// Echo implements the YourService/Echo RPC method
func (s *server) Echo(ctx context.Context, in *golang.StringMessage) (*golang.StringMessage, error) {
	// Just echo back the value from the request
	log.Printf("Received message: %s", in.GetValue())
	return &golang.StringMessage{Value: "Echo: " + in.GetValue()}, nil
}

func StartGRPCServer() {
	// Set up the listener on a specific port (e.g., 50051)
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the server with the gRPC server
	golang.RegisterYourServiceServer(s, &server{})

	// Start the server
	fmt.Println("Starting server on port :9090...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
