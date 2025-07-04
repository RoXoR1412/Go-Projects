package main

import (
	"context"
	"io"
	"log"
	pb "github.com/RoXoR1412/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Sreaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err ==io.EOF{
			break
		}
		if err !=nil{
			log.Fatalf("error receiving message: %v", err)
		}
		log.Println(message.Message)
	}
	log.Printf("Streaming Finished")
}