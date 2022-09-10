package main

import (
	"context"
	"log"

	"go-grpc/calculator/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c proto.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &proto.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Println(e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt of %d: %f\n", n, res.Result)
}
