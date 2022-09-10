package main

import (
	"context"
	"log"
	"time"

	"go-grpc/greet/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) GreetWithDeadline(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(greetWithDeadlineTime)
	}

	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
