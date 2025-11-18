package main

import (
	"io"
	"log"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	var res string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res += "Hello " + req.FirstName + "!\n"
	}
}
