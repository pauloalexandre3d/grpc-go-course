package main

import (
	"fmt"
	"time"

	pb "github.com/pauloalexandre3d/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s number %d", req.FirstName, i)
		res := &pb.GreetResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
