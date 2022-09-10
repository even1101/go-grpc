package main

import (
	"context"
	"go-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Println("greet %v", in)
	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
