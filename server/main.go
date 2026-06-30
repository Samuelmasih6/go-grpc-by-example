package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	grpcServer := grpc.NewServer()

	//Todo

	log.Println("server is running at ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to server: ", err)
	}

}
