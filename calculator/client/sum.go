package main

import (
	"context"
	"log"

	pb "github.com/pauloalexandre3d/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	response, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 120, SecondNumber: 80})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}
	log.Printf("Sum result: %d", response.Result)
}
