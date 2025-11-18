package main

import (
	"context"
	"log"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	response, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Paulo"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.Result)
}
