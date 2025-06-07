package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"protobuf/pb/chat"
	"protobuf/pb/user"
)

type userService struct {
	user.UnimplementedUserServiceServer
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {
	log.Printf("Received request to create user: %s", userRequest.FullName)
	return &user.CreateResponse{Message: "User created successfully"}, nil
}

type chatService struct {
	chat.UnimplementedChatServiceServer
}

func (cs *chatService) SendMessage(stream grpc.ClientStreamingServer[chat.ChatMessage, chat.ChatResponse]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Client has finished sending messages")
				break
			}
			return status.Errorf(codes.Unknown, "cannot receive a message: %v", err)
		}

		log.Printf("Received message: %s, to %d", req.Content, req.UserId)
	}
	return stream.SendAndClose(&chat.ChatResponse{
		Message: "Message received successfully",
	})
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})
	chat.RegisterChatServiceServer(serv, &chatService{})

	reflection.Register(serv)

	if err := serv.Serve(lis); err != nil {
		log.Fatal("Error serving:", err)
	} else {
		log.Println("Server started successfully on port 8080")
		defer lis.Close()
		defer serv.Stop()
		log.Println("Server stopped gracefully")
	}
}
