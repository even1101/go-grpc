package main

import (
	"io"
	"log"

	"go-grpc/calculator/proto"
)

func (*Server) Max(stream proto.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err = stream.Send(&proto.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}