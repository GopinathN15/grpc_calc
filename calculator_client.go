package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GopinathN15/grpc_calc"  // Update with your actual module path
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	req := &calculator.AddRequest{
		A: 3,
		B: 4,
	}

	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to add: %v", err)
	}

	fmt.Printf("Result: %d\n", res.Result)
}
