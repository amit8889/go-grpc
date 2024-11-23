package main

import (
	"context"
	"log"
	"time"

	pb "github.com/amit8889/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error in server starting : %v", err)
	}
	defer conn.Close()
	log.Println("Client Server started successfully")

	client := pb.NewGreeterServiceClient(conn)
	callSayHello(client)

}
func callSayHello(client pb.GreeterServiceClient) {

	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet : %v", err)
	}
	log.Printf("Greeting: %s", res.Message)

}
