package main

import (
	"go-grpc/greet/proto"
	"io"
	"log"
)

func (*Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %v\n", req)
		res += "Hello " + req.FirstName + "!\n"
	}
}
