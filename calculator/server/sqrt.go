package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"go-grpc/calculator/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) Sqrt(ctx context.Context, req *proto.SqrtRequest) (*proto.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with number %d\n", req.Number)

	number := req.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}

	return &proto.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
