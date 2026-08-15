package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Samuelmasih6/go-grpc-by-example/server/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:500051"

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect: ", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := pb.Addrequest{
		A: 10,
		B: 20,
	}
	res, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatalln("could not add: ", err)
	}
	log.Println("Sum: ", res.Sum)

}
