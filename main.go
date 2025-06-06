package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"protobuf/pb/user"
)

type userService struct {
	user.UnimplementedUserServiceServer
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {
	log.Printf("Received request to create user: %s", userRequest.FullName)
	return &user.CreateResponse{Message: "User created successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})

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
