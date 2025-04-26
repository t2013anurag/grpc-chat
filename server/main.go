package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-chat/chat/chat"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(streams pb.ChatService_ChatServer) error {
	for {
		msg, err := streams.Recv()
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}

		log.Printf("[%s]: %s", msg.User, msg.Message)

		// Echo back the message
		if err := streams.Send(&pb.ChatMessage{
			User:    "Server",
			Message: fmt.Sprintf("recevied: %s", msg.Message),
		}); err != nil {
			log.Printf("Error sending message: %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})

	log.Println("Server is running on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
