package main

import (
	"log"
	"net"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Listening on localhost:50051")

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, NewServer())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
