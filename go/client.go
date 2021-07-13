package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"felipeagger/gRPC/pb/data/protos"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8200", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	usr := protos.NewUserServiceClient(conn)

	fmt.Print("Enter the UserName: ")
	var inputUser string
	fmt.Scanln(&inputUser)

	response, err := usr.GetUser(context.Background(), &protos.UserRequest{Username: inputUser})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}

	log.Printf("Response from server: %v", response)
}
