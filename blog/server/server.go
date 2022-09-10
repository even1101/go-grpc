package main

import pb "go-grpc/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
