package main

import (
	"felipeagger/gRPC/pb/data/protos"
	"felipeagger/gRPC/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	usr := user.Server{}

	grpcServer := grpc.NewServer()

	protos.RegisterUserServiceServer(grpcServer, &usr)

	log.Println("Listening on Port: 8200!")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
