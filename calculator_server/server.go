package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/khoata39/grpccalculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.CalculatorServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

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

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
