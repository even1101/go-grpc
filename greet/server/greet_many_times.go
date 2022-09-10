package main

import (
	"fmt"
	"log"

	"go-grpc/greet/proto"
)

func (*Server) GreetManyTimes(in *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&proto.GreetResponse{
			Result: res,
		})
	}

	return nil
}
