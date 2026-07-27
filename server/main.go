package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Samuelmasih6/go-grpc-by-example/server/proto/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.Addrequest) (*pb.Addresponse, error) {
	sum := req.A + req.B
	log.Println("Sum: ", sum)
	return &pb.Addresponse{Sum: sum}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServer(grpcServer, &server{})

	log.Println("server is running at ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to server: ", err)
	}

}
