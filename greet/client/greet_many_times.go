package main

import (
	"context"
	"io"
	"log"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Paulo",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling GreetManyTimes: ", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while reading stream: ", err)
		}
		log.Println("Response from GreetManyTimes: ", msg.Result)
	}
}
