package main

import (
	"context"
	"log"

	pb "A4/user"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user := &pb.User{Id: 1, Name: "John Doe", Email: "john@example.com"}
	addedUser, err := client.AddUser(context.Background(), user)
	if err != nil {
		log.Fatalf("AddUser failed: %v", err)
	}
	log.Printf("User added: %v", addedUser)

	getUserResponse, err := client.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	log.Printf("User retrieved: %v", getUserResponse)

	listUsersResponse, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("ListUsers failed: %v", err)
	}
	log.Printf("Users list: %v", listUsersResponse.Users)
}
