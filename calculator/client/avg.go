package main

import (
	"context"
	"log"

	"go-grpc/calculator/proto"
)

func doAvg(c proto.CalculatorServiceClient) {
	log.Println("doAvg was invoked")
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)

		stream.Send(&proto.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
