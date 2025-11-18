package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: %w", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
	doGreetWithDeadline(c, 1*time.Second)
}
