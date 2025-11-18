package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	println("doLongGreet was invoked")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		panic(err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Paulo"},
		{FirstName: "Alexandre"},
		{FirstName: "Machado"},
	}

	for _, req := range reqs {
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v\n", err)
	}

	println(res.Result)
}
