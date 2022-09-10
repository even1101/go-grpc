package main

import (
	"go-grpc/greet/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("failed to listen on: %v ", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to server: g%v ", err)
	}
}
