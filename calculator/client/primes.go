package main

import (
	"context"
	"io"
	"log"

	"go-grpc/calculator/proto"
)

func doPrimes(c proto.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &proto.PrimeRequest{
		Number: 12390392840,
	}
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res.Result)
	}
}
