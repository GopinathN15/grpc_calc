package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/GopinathN15/grpc_calc"  // Update with your actual module path
	"google.golang.org/grpc"
)

type calculatorServer struct{}

func (s *calculatorServer) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	result := req.A + req.B
	return &calculator.AddResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	calculator.RegisterCalculatorServer(server, &calculatorServer{})

	fmt.Println("Server is listening on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
