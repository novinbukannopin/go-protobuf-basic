package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"protobuf/pb/user"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	userClient := user.NewUserServiceClient(conn)
	response, err := userClient.CreateUser(context.Background(), &user.User{
		Id:       1,
		FullName: "John Doe",
		Age:      13,
		Balance:  130000,
		Address: &user.Address{
			Id:          1,
			FullAddress: "123 Main St",
			City:        "Anytown",
			Province:    "Anystate",
		},
	})

	if err != nil {
		log.Fatal("Error creating user:", err)
	}
	log.Println("User created successfully", response.Message)
}
