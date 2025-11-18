package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Paulo"},
		{FirstName: "Alexandre"},
		{FirstName: "Machado"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending request: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			log.Printf("Received: %v", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
