package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/pauloalexandre3d/grpc-go-course/calculator/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: %w", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	doSum(c)

}
