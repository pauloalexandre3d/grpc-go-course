package main

import (
	"context"
	"log"

	pb "github.com/pauloalexandre3d/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(),
		&pb.SqrtRequest{
			Number: number,
		},
	)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error details from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatal("A non gRPC error occurred: ", err)
		}
	}

	log.Printf("Sqrt: %v", res.Result)
}
