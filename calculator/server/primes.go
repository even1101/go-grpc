package main

import (
	"log"

	"go-grpc/calculator/proto"
)

func (*Server) Primes(in *proto.PrimeRequest, stream proto.CalculatorService_PrimesServer) error {
	log.Printf("Primes was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&proto.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
