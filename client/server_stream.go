package main

import (
	"context"
	pb "github.com/olzzhas/go-grpc-akhilsharma/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while straming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
