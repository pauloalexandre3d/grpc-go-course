package main

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("The client canceled the request!")
			return nil, status.Errorf(codes.Canceled, "The client canceled the request!")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
