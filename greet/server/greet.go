package main

import (
	"context"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
