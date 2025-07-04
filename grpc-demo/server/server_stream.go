package main

import (
	"time"
	"log"
	pb "github.com/RoXoR1412/grpc-demo/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello" +name,
		}
		if err := stream.Send(res); err !=nil{
			return err
		}
		time.Sleep(2*time.Second)
	}
	return nil
}