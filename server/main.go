package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/amit8889/go-grpc/proto"
	"google.golang.org/grpc"
)

const (
	PORT = ":8080"
)

type helloServer struct {
	pb.GreeterServiceServer // Embedding this ensures forward compatibility
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Println("Greet request received from client")
	return &pb.HelloResponse{Message: "Hello from the server"}, nil
}

func main() {
	// Channel to handle OS signals for graceful shutdown.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT)

	go func() {
		// Start listening on the specified port.
		lis, err := net.Listen("tcp", PORT)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
		defer lis.Close()

		// Create and start the gRPC server.
		grpcServer := grpc.NewServer()
		pb.RegisterGreeterServiceServer(grpcServer, &helloServer{}) // Pass properly instantiated helloServer
		log.Printf("Server started at: %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error in gRPC server: %v", err)
		}
	}()

	// Wait for an interrupt signal.
	<-done
	log.Println("Server is shutting down...")

	// Graceful shutdown context.
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Server has shut down gracefully.")
}
