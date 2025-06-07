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
	stream, err := chatClient.Chat(context.Background())
	if err != nil {
		log.Fatal("Error creating chat stream:", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  1,
		Content: "Hello this is client",
	})

	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	msg, err := stream.Recv()

	if err != nil {
		log.Fatal("Error receiving message:", err)
	}

	log.Printf("Received message from user %d: %s", msg.UserId, msg.Content)

	err = stream.Send(&chat.ChatMessage{
		UserId:  2,
		Content: "Hello this is client agaoin",
	})
	if err != nil {
		log.Fatal("Error sending message:", err)
	}

	msg, err = stream.Recv()
	if err != nil {
		log.Fatal("Error receiving message:", err)
	}

	log.Printf("Received message from user %d: %s", msg.UserId, msg.Content)

	//for {
	//	msg, err := stream.Recv()
	//	if err != nil {
	//		if errors.Is(err, io.EOF) {
	//			log.Println("No more messages from server")
	//			break
	//		}
	//		log.Fatal("Error receiving message:", err)
	//	}
	//	log.Printf("Received message from user %d: %s", msg.UserId, msg.Content)
	//}

	//err = stream.Send(&chat.ChatMessage{
	//	UserId:  1,
	//	Content: "hello world",
	//})
	//
	//if err != nil {
	//	log.Fatal("Error sending message:", err)
	//}
	//
	//err = stream.Send(&chat.ChatMessage{
	//	UserId:  2,
	//	Content: "hello golang",
	//})
	//
	//if err != nil {
	//	log.Fatal("Error sending message:", err)
	//}
	//
	//res, err := stream.CloseAndRecv()
	//if err != nil {
	//	log.Fatal("Error receiving response:", err)
	//}
	//log.Println("Connected to chat service successfully", res.Message)
}
