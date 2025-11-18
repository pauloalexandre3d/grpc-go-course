package main

import (
	"context"

	pb "github.com/pauloalexandre3d/grpc-go-course/calculator/proto"
)

type Server struct {
	pb.CalculatorServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}
