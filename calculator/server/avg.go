package main

import (
	"io"
	"log"

	"go-grpc/calculator/proto"
)

func (*Server) Avg(stream proto.CalculatorService_AvgServer) error {
	log.Println("Avg was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sum += req.Number
		count++
	}
}
