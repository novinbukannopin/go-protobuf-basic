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
	"protobuf/pb/common"
	"protobuf/pb/user"
)

type userService struct {
	user.UnimplementedUserServiceServer
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {

	if userRequest.Age < 1 {
		return &user.CreateResponse{
			BaseResponse: &common.BaseResponse{
				StatusCode: 400,
				IsSuccess:  false,
				Message:    "Bad Request",
				ValidationErrors: []*common.ValidationError{
					{
						Field:   "age",
						Message: "Age must be greater than 0",
					},
					{
						Field:   "full_name",
						Message: "Full name is required",
					},
				},
			},
		}, nil
	}

	return &user.CreateResponse{
		BaseResponse: &common.BaseResponse{
			StatusCode: 200,
			IsSuccess:  true,
			Message:    "User created successfully",
		},
	}, nil
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

func (cs *chatService) ReceiveMessages(req *chat.ReceiveMessageRequest, stream grpc.ServerStreamingServer[chat.ChatMessage]) error {
	log.Printf("Received request to receive messages from user: %s", req.UserId)

	for i := 0; i < 5; i++ {
		err := stream.Send(&chat.ChatMessage{
			UserId:  req.UserId,
			Content: "Message " + string(rune(i+1)),
		})

		if err != nil {
			return status.Errorf(codes.Unknown, "cannot send message: %v", err)
		}
	}

	return nil
}

func (cs *chatService) Chat(stream grpc.BidiStreamingServer[chat.ChatMessage, chat.ChatMessage]) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Client has finished sending messages")
				break
			}
			return status.Errorf(codes.Unknown, "cannot receive a message: %v", err)
		}
		log.Printf("Received message from user %d: %s", msg.UserId, msg.Content)

		err = stream.Send(&chat.ChatMessage{
			UserId:  msg.UserId,
			Content: "reply to server",
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot send a message: %v", err)
		}
	}

	return nil
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
