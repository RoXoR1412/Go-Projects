package main

import (
	"time"
	"log"
	pb "github.com/RoXoR1412/grpc-demo/proto"
	"context"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Client Streaming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("error sending name: %v", err)
		}
		log.Printf("Sent the request with name : %s", name)
		time.Sleep(2*time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished")
	if err != nil{
		log.Fatalf("Error while receiving %v",err)
	}
	log.Printf("Response from server: %s", res)
}