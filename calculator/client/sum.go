package main

import (
	"context"
	"log"

	"go-grpc/calculator/proto"
)

func doSum(c proto.CalculatorServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Sum(context.Background(), &proto.SumRequest{FirstNumber: 10, SecondNumber: 20})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Total: %d\n", r.Result)
}
