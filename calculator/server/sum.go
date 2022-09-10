package main

import (
	"context"
	"log"

	"go-grpc/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Sum was invoked with %v\n", in)
	return &proto.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
