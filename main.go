package main

import (
	"context"
	pb "github.com/quabynah-bilson/grpc-go-client/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.Println("this is my grpc go client")
	
	// create a connection to server
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	
	// create echo service client
	client := pb.NewEchoServiceClient(conn)
	
	ctx := context.Background()
	req := &pb.EchoRequest{Name: "Quabynah Bilson"}
	res, err := client.EchoMessage(ctx, req)
	if err != nil {
		log.Fatalf("there was an error getting the response: %v", err)
	}
	
	log.Println(res.GetGreeting())
}
