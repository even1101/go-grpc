package main

import (
	"context"
	"log"

	"go-grpc/greet/proto"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Greet(context.Background(), &proto.GreetRequest{FirstName: "Clement"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
