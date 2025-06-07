package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"protobuf/pb/chat"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	//userClient := user.NewUserServiceClient(conn)
	//response, err := userClient.CreateUser(context.Background(), &user.User{
	//	Id:       1,
	//	FullName: "John Doe",
	//	Age:      13,
	//	Balance:  130000,
	//	Address: &user.Address{
	//		Id:          1,
	//		FullAddress: "123 Main St",
	//		City:        "Anytown",
	//		Province:    "Anystate",
	//	},
	//})
	//
	//if err != nil {
	//	log.Fatal("Error creating user:", err)
	//}
	//log.Println("User created successfully", response.Message)

	chatClient := chat.NewChatServiceClient(conn)
	stream, err := chatClient.SendMessage(context.Background())
	if err != nil {
		log.Fatal("Error creating chat stream:", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  1,
		Content: "hello world",
	})

	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  2,
		Content: "hello golang",
	})

	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error receiving response:", err)
	}
	log.Println("Connected to chat service successfully", res.Message)
}
