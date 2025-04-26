package main

import (
	"bufio"
	"context"
	"google.golang.org/grpc"
	pb "grpc-chat/chat/chat"
	"log"
	"os"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Failed to start chat: %v", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving: %v", err)
			}
			log.Printf("[Server]: %s", resp.Message)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		log.Print("Enter message: ")
		text, _ := reader.ReadString('\n')

		if err := stream.Send(&pb.ChatMessage{
			User:    "Client",
			Message: text,
		}); err != nil {
			log.Fatalf("Error sending: %v", err)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
