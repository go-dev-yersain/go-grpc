package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/simplegrpc/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)

	x := req.GetFirstNumber()
	y := req.GetSecondNumber()
	sum := x + y
	res := &calculatorpb.SumResponse{
		Result: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello calc")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, new(server))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
