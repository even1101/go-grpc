package main

import (
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
}

func main() {
	_, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("failed to listen on: %v ", err)
	}
}
