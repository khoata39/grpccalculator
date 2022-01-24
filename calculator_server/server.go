package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "../calulatorpb/calculator.pb.go"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("Sum function called with : %v\n", req)

	firstNumber := req.GetFirst_Number()
	secondNumber := req.GetSecond_Number()
	sum := firstNumber + secondNumber

	return &pb.SumResponse{
		Result: sum,
	}, nil
}

func main() {
	fmt.Println("Hello, I am gRPC Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
