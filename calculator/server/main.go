package main

import (
	"go-grpc/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("failed to listen on: %v ", err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to server: g%v ", err)
	}
}
